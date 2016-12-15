package main

// Microservice
// Storage
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//import "fmt"

// NewStorage - create a new Storage
func NewStorage(conf *Tuner) *Storage {
	logger := *NewLogger()
	metric := *NewMetric(&logger)
	s := &Storage{
		Logger: logger,
		Metric: metric,
	}
	return s
}

/*
Storage structure
It creates and stores objects
*/
type Storage struct {
	Logger Logger
	Metric Metric
}
