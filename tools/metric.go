package tools

// Microservice
// Metric
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

// Metric - an example of a middleware that logged the duration of the application
func Metric(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	logger := logrus.New()
	return func(w http.ResponseWriter, r *http.Request) {
		timeStart := int(time.Now().UnixNano())
		f(w, r)
		go logger.WithField("duration", int(time.Now().UnixNano())-timeStart).
			WithField("url", r.URL.Path).
			Print()
	}
}
