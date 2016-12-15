// Microservice
// Handler
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package main

import "net/http"

// NewHandler - create a new Handler
func NewHandler(conf *Tuner) *Handler {
	h := &Handler{
		Store: NewStorage(conf),
		Conf:  conf,
	}
	return h
}

// Handler structure
type Handler struct {
	Store *Storage
	Conf  *Tuner
}

// Handle - add the handle to the handler
func (h *Handler) Handle(f func(http.ResponseWriter, *http.Request)) *Handle {
	x := NewHandle(f)
	return x
}

// Test - handler method for example
func (h *Handler) Test(w http.ResponseWriter, req *http.Request) {
	w.Header().Del("Content-Type")
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	// fmt.Print("Testus!\n")
	w.Write([]byte("Test"))
}
