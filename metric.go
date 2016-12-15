// Microservice
// Metric
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package main

import (
	"context"
	//"fmt"
	"net/http"
	"time"
)

// NewMetric - create a new Metric
func NewMetric(logger *Logger) *Metric {
	m := &Metric{logger: logger}
	return m
}

// Metric structure
// This library shows a simple version of the logging duration metrics
type Metric struct {
	logger *Logger
}

// Start - fix a starting time
func (g *Metric) Start(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, bool) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "timeStart", int(time.Now().UnixNano()))
	req = req.WithContext(ctx)
	return w, req, true
}

// End - sending metrics on the duration
func (g *Metric) End(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, bool) {
	timeStart := req.Context().Value("timeStart").(int)
	go g.send(map[string]interface{}{"duration": int(time.Now().UnixNano()) - timeStart})
	return w, req, true
}

func (m *Metric) send(fields map[string]interface{}) {
	msg := m.logger.Message()
	for k, v := range fields {
		msg.Field(k, v)
	}
	msg.Send()
}
