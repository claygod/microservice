package tools

// Microservice tools
// Test Metric
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"testing"
)

func TestMetricStart(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	m := NewMetric()
	_, req = m.Start(nil, req)
	if req == nil {
		t.Error("Error in metric `Start`")
	}
	if req.Context().Value("timeStart") == nil {
		t.Error("Error while creating the parameter start time `End`")
	}
}

func TestMetricEnd(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	m := NewMetric()
	_, req = m.Start(nil, req)
	_, req = m.End(nil, req)
	if req == nil {
		t.Error("Error in metric `End`")
	}
}

func TestMetricStartNil(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	m := NewMetric()
	_, req = m.Start(nil, nil)
	if req != nil {
		t.Error("Transmission error `nil` as request")
	}
}
