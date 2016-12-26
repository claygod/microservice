package main

// Microservice
// Handler
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "fmt"
import "net/http"
import "github.com/Sirupsen/logrus"

// NewHandler - create a new Handler
func NewHandler(conf *Tuner) *Handler {
	h := &Handler{
		Conf: conf,
		Log:  logrus.New(),
	}
	return h
}

// Handler structure
type Handler struct {
	Conf *Tuner
	Log  *logrus.Logger
}

// Queue( - create the queue to the handler
func (h *Handler) Queue(args ...func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)) *Queue {
	x := NewQueue(args)
	return x
}

// HelloWorld - handler method for example
func (h *Handler) HelloWorld(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
	if req != nil {
		w.Header().Del("Content-Type")
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.Write([]byte(fmt.Sprintf("Hello %s!", h.Conf.Main.Name)))
		h.Log.Infof("log=%s hello=%s", "demo", "world")
	}
	return w, req
}
