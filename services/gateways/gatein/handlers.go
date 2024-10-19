package gatein

// Microservice
// Gate In (handlers)
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/claygod/microservice/usecases"
	"github.com/julienschmidt/httprouter"
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

	result, err := g.foobarInteractor.GetBar(params.ByName("key"), ctx)

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

func (g *GateIn) HealthCheckHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

func (g *GateIn) ReadynessHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if g.hasp.IsReady() || g.hasp.IsRun() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

func (g *GateIn) Metrics(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	g.metrics.Handler().ServeHTTP(w, req)
}

func (g *GateIn) writeError(lg *slog.Logger, w http.ResponseWriter, err error, status int) {
	go g.metrics.ResponceCode(status)
	go lg.Error(err.Error())

	http.Error(w, err.Error(), status)
}
