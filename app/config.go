package app

// Microservice
// Config
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/microservice/services/gateways/bar"
	"github.com/claygod/microservice/services/gateways/gatein"
	"github.com/claygod/microservice/services/repositories/foo"
	"github.com/claygod/microservice/usecases"
)

type Config struct {
	Dummy      string
	Interactor *usecases.Config
	GateIn     *gatein.Config
	FooRepo    *foo.Config
	BarGate    *bar.Config
}
