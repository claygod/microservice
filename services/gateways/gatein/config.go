package gatein

// Microservice
// Gate In (config)
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Config struct {
	Title string
	Port  string
}

const (
	headerRequestID = "LogID"

	nanoToMilli     = 1000000
	timeStopService = time.Second * 15
)

func (g *GateIn) getReqID(req *http.Request) string {
	x := req.Header.Get(headerRequestID)

	if x == "" {
		x = fmt.Sprintf("%s-%s", g.config.Title, uuid.New().String())
	}

	return x
}

func (g *GateIn) checkPanic(req *http.Request) {
	if r := recover(); r != nil {
		g.logger.Error(fmt.Sprintf("PANIC (GateIn):Recovered: %v %s: %s", r,
			headerRequestID, g.getReqID(req)))
	}
}
