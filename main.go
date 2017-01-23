// This library provides a simple framework of microservice, which includes
// a configurator, a metrics, and of course the handler
package main

// Microservice
// Main
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import "github.com/claygod/BxogV2"
import "github.com/claygod/microservice/tools"

// Main
func main() {
	conf, err := NewTuner("config.toml")
	if err != nil {
		panic(err)
	}

	hr := NewHandler(conf)

	m := bxogv2.New()
	m.Add("/", tools.Metric(hr.HelloWorld))
	m.Start(":" + conf.Main.Port)
}
