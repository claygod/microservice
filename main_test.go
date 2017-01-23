package main

// Microservice
// Test Main
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"net/http/httptest"
	"testing"

	bx "github.com/claygod/BxogV2"
)

func TestMain(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	hello :=
		func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(777)
			return
		}
	m := bx.New()
	m.Add("/", hello)
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
	f := func(w http.ResponseWriter, req *http.Request) {
		return
	}

	m := bx.New()
	m.Add("/", f)
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
	f := func(w http.ResponseWriter, req *http.Request) {
		return
	}

	m := bx.New()
	m.Add("/", f)
	m.Test()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.ServeHTTP(w, req)
		}
	})
}
