package main

// Microservice
// Queue
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "net/http"

// NewQueue - create a new Queue
func NewQueue(args []func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)) *Queue {
	q := &Queue{}
	for _, f := range args {
		q.list = append(q.list, f)
	}
	return q
}

// Queue structure
type Queue struct {
	list []func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)
}

// Run - launch handler's queue
func (q *Queue) Run(w http.ResponseWriter, req *http.Request) {
	for _, f := range q.list {
		if w, req = f(w, req); req == nil {
			return
		}
	}
}
