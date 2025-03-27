package gatein

// Microservice
// Gate In
// Copyright © 2021-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"

	"github.com/claygod/microservice/domain"
	"github.com/claygod/microservice/services/metrics"
	"github.com/claygod/microservice/usecases"
)

const (
	reqTimeout       = 30 * time.Second
	swagYAMLFileName = "swagger.yaml"
	emptyString      = ""
)

type GateIn struct {
	hasp             domain.StartStopInterface
	logger           *slog.Logger
	router           *httprouter.Router
	server           *http.Server
	foobarInteractor *usecases.FooBarInteractor
	config           *Config
	metrics          *metrics.Metrics
	swagAPI          *swagger.API
	swagGenerated    bool
}

func New(ss domain.StartStopInterface, lg *slog.Logger, cnf *Config, fbi *usecases.FooBarInteractor, mtr *metrics.Metrics) *GateIn {
	g := &GateIn{
		logger:           lg,
		foobarInteractor: fbi,
		config:           cnf,
		hasp:             ss,
		metrics:          mtr,
	}

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	// optional route
	router.GET("/", g.middle(g.WelcomeHandler))

	// public routes
	router.GET("/public/v1/bar/:key", g.middle(g.GetBarHandler))
	swagV1Bar := endpoint.New("get", "/public/v1/bar/{key}", "Find object bar by key",
		endpoint.Handler(g.GetBarHandler),
		endpoint.Path("key", "string", "key of object bar to return", true),
		endpoint.Response(http.StatusOK, domain.Bar{Data: "three"}, "successful operation"),
		endpoint.Response(http.StatusNotFound, emptyString, "not found"),
		endpoint.Response(http.StatusInternalServerError, emptyString, "internal server error"),
		// endpoint.Tags("section A"),
	)

	// private routes
	// example: /private/v1/foo/:key

	api := swag.New(
		swag.Description(cnf.Title+" API"),
		swag.Version("1.0"),
		swag.Title(cnf.Title+" swagger"),
		swag.ContactEmail("mail@mail.com"),
		swag.BasePath("/public"),
		swag.Endpoints(swagV1Bar),
		// swag.Tag("section A", "this is section A", swag.TagURL("")),
	)

	g.swagAPI = api

	// service routes
	router.GET("/healthz/ready", g.middle(g.HealthCheckHandler)) // for SRE
	router.GET("/healthz", g.middle(g.HealthCheckHandler))       // for kubernetes
	router.GET("/readyness", g.middle(g.ReadynessHandler))       // for kubernetes
	router.GET("/metrics", g.Metrics)
	router.GET("/swagger", g.middle(g.SwaggerHandler))

	g.router = router

	return g
}

func (g *GateIn) Start() error {
	if !g.hasp.Start() {
		return errors.New("gatein:failed to start")
	}

	g.server = &http.Server{
		Handler:      g.router,
		Addr:         g.config.Port,
		WriteTimeout: reqTimeout,
		ReadTimeout:  reqTimeout,
	}

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			g.logger.Error(fmt.Sprintf("Httpserver: ListenAndServe() error: %s", err))
		} else {
			g.logger.Info("Httpserver: ListenAndServe() closing...")
		}
	}()

	return nil
}

func (g *GateIn) Stop() error {
	if !g.hasp.Stop() {
		return errors.New("gatein:failed to stop")
	}

	return nil
}

func (g *GateIn) middle(f func(http.ResponseWriter, *http.Request, httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		timeStart := time.Now().UnixNano()

		if !g.hasp.IsRun() {
			http.Error(w, "service is stopped", http.StatusServiceUnavailable)

			return
		}

		g.hasp.Add()
		defer g.hasp.Done()

		defer g.checkPanic(req)

		f(w, req, p)

		if reqURI := req.URL.RequestURI(); reqURI != "/health_check" && reqURI != "/readyness" && reqURI != "/metrics" {
			durNano := (time.Now().UnixNano() - timeStart)
			g.metrics.Request(req.URL.RequestURI(), time.Duration(durNano))

			dur := durNano / nanoToMilli
			go g.logger.With(headerRequestID, g.getReqID(req)).Info(fmt.Sprintf("duration: %d ms , link: %s", dur, req.URL.RequestURI()))
		}
	}
}
