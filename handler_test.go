package main

// Microservice
// Test Handler
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHelloWorld(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	conf, err := NewTuner("config.toml")
	if err != nil {
		panic(err)
	}
	hr := NewHandler(conf)
	hr.HelloWorld(w, req)
	if req.Header.Get("Test") != "Test" {
		t.Error("Error Handler in the `HelloWorld`")
	}
}
