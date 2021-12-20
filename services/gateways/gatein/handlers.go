package gatein

// Microservice
// Gate In (handlers)
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/claygod/microservice/usecases"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (g *GateIn) WelcomeHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if _, err := io.WriteString(w, "Welcome to Service!..."); err != nil {
		lg := g.logger.WithField(headerRequestID, g.getReqID(req))

		g.writeError(lg, w, err, http.StatusInternalServerError)

		return
	}

	go g.metrics.ResponceCode(http.StatusOK)
}

func (g *GateIn) GetBarHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx, _ := context.WithTimeout(req.Context(), 30*time.Second)
	lg := g.logger.WithField(headerRequestID, g.getReqID(req))

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
	w.Write([]byte(out))
}

func (g *GateIn) HealthCheckHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	lg := g.logger.WithField(headerRequestID, g.getReqID(req))

	out, err := json.Marshal(g.foobarInteractor.GetHealth())

	if err != nil {
		g.writeError(lg, w, err, http.StatusInternalServerError)

		return
	}

	go g.metrics.ResponceCode(http.StatusOK)
	w.Write([]byte(out))
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

func (g *GateIn) writeError(lg *logrus.Entry, w http.ResponseWriter, err error, status int) {
	go g.metrics.ResponceCode(status)
	go lg.Error(err)

	http.Error(w, err.Error(), status)
}
