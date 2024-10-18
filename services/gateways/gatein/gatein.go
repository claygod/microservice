package gatein

// Microservice
// Gate In
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"github.com/claygod/microservice/domain"
	"github.com/claygod/microservice/services/metrics"
	"github.com/claygod/microservice/usecases"
)

type GateIn struct {
	hasp             domain.StartStopInterface
	logger           *logrus.Entry
	router           *httprouter.Router
	server           *http.Server
	foobarInteractor *usecases.FooBarInteractor
	config           *Config
	metrics          *metrics.Metrics
}

func New(ss domain.StartStopInterface, lg *logrus.Entry, cnf *Config,
	fbi *usecases.FooBarInteractor, mtr *metrics.Metrics) *GateIn {
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

	// service routes
	router.GET("/health_check", g.middle(g.HealthCheckHandler)) // for SRE
	router.GET("/readyness", g.middle(g.ReadynessHandler))      // for kubernetes
	router.GET("/metrics", g.Metrics)

	// public routes
	router.GET("/piblic/v1/bar/:key", g.middle(g.GetBarHandler))

	// private routes
	// example: /private/v1/foo/:key

	g.router = router

	return g
}

func (g *GateIn) Start() error {
	if !g.hasp.Start() {
		return errors.New("gatein:failed to start")
	}

	g.server = &http.Server{
		Handler: g.router,
		Addr:    g.config.Port,
		// Good practice: enforce timeouts for servers you create!
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			g.logger.Errorf("Httpserver: ListenAndServe() error: %s", err)
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

	// ctx, cancel := context.WithTimeout(context.Background(), timeStopService)

	// defer cancel()

	// g.logger.Info(g.server.Shutdown(ctx)) // failure/timeout shutting down the server gracefully

	return nil
}

func (g *GateIn) middle(f func(http.ResponseWriter, *http.Request,
	httprouter.Params)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
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
			go g.logger.WithField(headerRequestID, g.getReqID(req)).Info(fmt.Sprintf("duration: %d ms , link: %s", dur, req.URL.RequestURI()))
		}
	}
}
