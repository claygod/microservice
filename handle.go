// Microservice
// Handle
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

package main

import "net/http"

// NewHandle - create a new Handle
func NewHandle(f func(http.ResponseWriter, *http.Request)) *Handle {
	h := &Handle{executor: f}
	return h
}

// Handle structure
type Handle struct {
	executor     func(http.ResponseWriter, *http.Request)
	quequeBefore []func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)
	quequeAfter  []func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)
}

// Before - functions that run before the handler
func (h *Handle) Before(args ...func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)) *Handle {
	for _, f := range args {
		h.quequeBefore = append(h.quequeBefore, f)
	}
	return h
}

// After - functions that run after the handler
func (h *Handle) After(args ...func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)) *Handle {
	for _, f := range args {
		h.quequeAfter = append(h.quequeAfter, f)
	}
	return h
}

// Run - launch handler
func (h *Handle) Run(w http.ResponseWriter, req *http.Request) {
	for _, f := range h.quequeBefore {
		if w2, req2, result := f(w, req); result {
			w = w2
			req = req2
		} else {
			return
		}
	}
	if h.executor != nil {
		h.executor(w, req)
	}
	for _, f := range h.quequeAfter {
		if w2, req2, result := f(w, req); result {
			w = w2
			req = req2
		} else {
			return
		}
	}
}
