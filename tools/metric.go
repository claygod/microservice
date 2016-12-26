package tools

// Microservice
// Metric
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"context"
	"net/http"
	"time"
)

// NewMetric - create a new Metric
func NewMetric() *Metric {
	m := &Metric{logger: NewLogger()}
	return m
}

// Metric structure
// This library shows a simple version of the logging duration metrics
type Metric struct {
	logger *Logger
}

// Start - fix a starting time
func (m *Metric) Start(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
	if req != nil {
		ctx := req.Context()
		ctx = context.WithValue(ctx, "timeStart", int(time.Now().UnixNano()))
		req = req.WithContext(ctx)
	}
	return w, req
}

// End - sending metrics on the duration
func (m *Metric) End(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
	if req != nil {
		timeStart := req.Context().Value("timeStart").(int)
		go m.send(map[string]interface{}{"duration": int(time.Now().UnixNano()) - timeStart})
	}
	return w, req
}

func (m *Metric) send(fields map[string]interface{}) {
	msg := m.logger.Message()
	for k, v := range fields {
		msg.Field(k, v)
	}
	msg.Send()
}
