package gatein

// Microservice
// Gate In (handlers)
// Copyright © 2021-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/claygod/microservice/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"sigs.k8s.io/yaml"
)

const (
	filePerm = fs.FileMode(0x644)
)

func (g *GateIn) WelcomeHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if _, err := io.WriteString(w, "Welcome to service "+g.config.Title); err != nil {
		lg := g.logger.With(headerRequestID, g.getReqID(req))

		g.writeError(lg, w, err, http.StatusInternalServerError)

		return
	}

	go g.metrics.ResponceCode(http.StatusOK)
}

func (g *GateIn) GetBarHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx, cancel := context.WithTimeout(req.Context(), reqTimeout)
	defer cancel()

	lg := g.logger.With(headerRequestID, g.getReqID(req))
	key := params.ByName("key")
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Var(key, "alphanum,gte=0,lte=30"); err != nil {
		g.writeError(lg, w, err, http.StatusBadRequest)

		return
	}

	result, err := g.foobarInteractor.GetBar(ctx, key)

	if err != nil {
		if errors.Is(err, usecases.ErrUserBadRequest) {
			g.writeError(lg, w, err, http.StatusBadRequest)
		} else {
			g.writeError(lg, w, err, http.StatusInternalServerError)
		}

		return
	} else if result == nil {
		g.writeError(lg, w, errors.New("not found"), http.StatusNotFound)

		return
	}

	out, err := json.Marshal(result)
	if err != nil {
		g.writeError(lg, w, err, http.StatusInternalServerError)

		return
	}

	go g.metrics.ResponceCode(http.StatusOK)
	if _, err := w.Write(out); err != nil {
		g.logger.Error(err.Error())
	}
}

func (g *GateIn) HealthCheckHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	lg := g.logger.With(headerRequestID, g.getReqID(req))

	out, err := json.Marshal(g.foobarInteractor.GetHealth())
	if err != nil {
		g.writeError(lg, w, err, http.StatusInternalServerError)

		return
	}

	go g.metrics.ResponceCode(http.StatusOK)
	if _, err := w.Write(out); err != nil {
		g.logger.Error(err.Error())
	}
}

func (g *GateIn) ReadynessHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if g.hasp.IsReady() || g.hasp.IsRun() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func (g *GateIn) Metrics(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	g.metrics.Handler().ServeHTTP(w, req)
}

func (g *GateIn) SwaggerHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	g.swagAPI.Host = req.Host
	g.swagAPI.Schemes[0] = g.schemeFromProto(req.Proto)

	apij, err := json.Marshal(g.swagAPI)
	if err != nil {
		g.logger.Error(err.Error())
		g.writeError(g.logger, w, err, http.StatusInternalServerError)
	} else {
		apij = []byte(strings.ReplaceAll(string(apij), `"tags":null,`, ""))

		if !g.swagGenerated {
			if err := g.saveSwaggerToFile(apij); err != nil {
				g.logger.Error(err.Error())
			} else {
				g.swagGenerated = true
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		if _, err := w.Write(apij); err != nil {
			g.logger.Error(err.Error())
		}
	}
}

func (g *GateIn) schemeFromProto(proto string) string {
	if strings.Contains(proto, "HTTPS") {
		return "https"
	}
	return "http"
}

func (g *GateIn) saveSwaggerToFile(apij []byte) error {
	y, err := yaml.JSONToYAML(apij)
	if err != nil {
		return err
	}

	fullPath := g.config.ConfigPath + "/" + swagYAMLFileName

	return os.WriteFile(fullPath, y, filePerm)
}

func (g *GateIn) writeError(lg *slog.Logger, w http.ResponseWriter, err error, status int) {
	go g.metrics.ResponceCode(status)
	go lg.Error(err.Error())

	http.Error(w, err.Error(), status)
}
