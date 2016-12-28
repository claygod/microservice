package main

// Microservice
// Test Main
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/claygod/Bxog"
)

func TestMain(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	hr := NewHandler(&Tuner{})

	hello := hr.Queue(
		func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
			w.WriteHeader(777)
			return w, req
		},
	)
	m := bxog.New()
	m.Add("/", hello.Run)
	m.Test()
	m.ServeHTTP(w, req)
	if w.Code != 777 {
		t.Error("Error application")
	}
}

func BenchmarkMain(b *testing.B) {
	b.StopTimer()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	q := NewQueue([]func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request){func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
		return w, req
	}})
	q.Run(w, req)

	m := bxog.New()
	m.Add("/", q.Run)
	m.Test()
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		m.ServeHTTP(w, req)
	}
}

func BenchmarkMainParallel(b *testing.B) {
	b.StopTimer()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	q := NewQueue([]func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request){func(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
		return w, req
	}})
	q.Run(w, req)

	m := bxog.New()
	m.Add("/", q.Run)
	m.Test()
	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.ServeHTTP(w, req)
		}
	})
}
