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
)

func (g *GateIn) WelcomeHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if _, err := io.WriteString(w, "Welcome to Service!..."); err != nil {
		go g.logger.WithField(headerRequestID, g.getReqID(req)).Error(err)
	}
}

func (g *GateIn) GetBarHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx, _ := context.WithTimeout(req.Context(), 30*time.Second)
	lg := g.logger.WithField(headerRequestID, g.getReqID(req))

	result, err := g.foobarInteractor.GetBar(params.ByName("key"), ctx)

	if err != nil {
		go lg.Error(err)

		if errors.Is(err, usecases.ErrUserBadRequest) {
			http.Error(w, err.Error(), http.StatusBadRequest)

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return

	} else if result == nil {
		http.Error(w, "not found", http.StatusNotFound)

		return
	}

	out, err := json.Marshal(result)

	if err != nil {
		go lg.Error(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte(out))
}

func (g *GateIn) HealthCheckHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	lg := g.logger.WithField(headerRequestID, g.getReqID(req))

	out, err := json.Marshal(g.foobarInteractor.GetHealth())

	if err != nil {
		go lg.Error(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

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
