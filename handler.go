package main

// Microservice
// Handler
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "fmt"
import "net/http"
import "github.com/Sirupsen/logrus"

//import "context"

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

// HelloWorld - handler method for example
func (h *Handler) HelloWorld(w http.ResponseWriter, req *http.Request) {
	if req != nil {
		w.Header().Del("Content-Type")
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.Write([]byte(fmt.Sprintf("Hello %s!", h.Conf.Main.Name)))
		// for tests
		req.Header.Set("Test", "Test")
		///ctx := req.Context()
		///ctx = context.WithValue(ctx, "Test", "Test")
		///req = req.WithContext(ctx)
		// demo log
		// go h.Log.WithField("hello", h.Conf.Main.Name).Error("Demo of error")
	}
	return
}
