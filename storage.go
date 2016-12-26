package main

// Microservice
// Storage
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "github.com/claygod/microservice/tools"

// NewStorage - create a new Storage
func NewStorage(conf *Tuner) *Storage {
	metric := *tools.NewMetric()
	session := *tools.NewSession(conf.Session.Secure, conf.Session.Name)
	s := &Storage{
		Metric:  metric,
		Session: session,
	}
	return s
}

// Storage structure
// It creates and stores objects
type Storage struct {
	Metric  tools.Metric
	Session tools.Session
}
