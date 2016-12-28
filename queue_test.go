package main

// Microservice
// Test Queue
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueue(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	f1 := func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
		return w, nil
	}
	f2 := func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
		w.WriteHeader(1000)
		return w, nil
	}
	q := NewQueue([]func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request){f1, f2})
	q.Run(w, req)

	if int(w.Code) == 1000 {
		t.Error("Error in Queue")
	}
}
