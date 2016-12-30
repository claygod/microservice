package main

// Microservice
// Test Handler
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerQueueWithArgs(t *testing.T) {
	hr := &Handler{}
	f1 := func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) { return nil, nil }
	f2 := func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) { return nil, nil }
	q := hr.Queue(f1, f2)
	if q == nil {
		t.Error("Failed to create a new queue")
	}
}

func TestHandlerQueueEmptyArgs(t *testing.T) {
	hr := &Handler{}
	q := hr.Queue()
	if q == nil {
		t.Error("Failed to create a new queue (without args)")
	}
}

func TestHandlerHelloWorld(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	conf, err := NewTuner("config.toml")
	if err != nil {
		panic(err)
	}
	hr := NewHandler(conf)
	_, req = hr.HelloWorld(w, req)
	if v := req.Context().Value("Test"); v == nil || v.(string) != "Test" {
		t.Error("Error Handler in the `HelloWorld`")
	}
}
