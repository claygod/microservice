//This library provides a simple framework of microservice, which includes
//a configurator, a logger, metrics, and of course the handler
package main

// Microservice
// Main
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "github.com/claygod/Bxog"

// Main
func main() {
	conf, err := NewTuner("config.toml")
	if err != nil {
		panic(err)
	}
	store := NewStorage(conf)

	hr := NewHandler(conf)

	hello := hr.Queue(
		store.Metric.Start,
		store.Session.Check,
		hr.HelloWorld,
		store.Metric.End,
	)

	m := bxog.New()
	m.Add("/", hello.Run)
	m.Start(":" + conf.Main.Port)
}
