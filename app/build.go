package app

// Microservice
// App build
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"time"

	"github.com/claygod/tools/startstop"
	"github.com/sirupsen/logrus"

	"github.com/claygod/microservice/services/gateways/bar"
	"github.com/claygod/microservice/services/gateways/gatein"
	"github.com/claygod/microservice/services/metrics"
	"github.com/claygod/microservice/services/repositories/foo"
	"github.com/claygod/microservice/usecases"
)

func New(cnf *Config) (*Application, error) {
	fooRepo := foo.New(startstop.New(1*time.Second), logrus.New().WithField("service", cnf.FooRepo.Title), cnf.FooRepo)
	gateBar := bar.New(startstop.New(1*time.Second), logrus.New().WithField("service", cnf.GateBar.Title), cnf.GateBar)
	fbi := usecases.NewFooBarInteractor(startstop.New(1*time.Second), cnf.Interactor, fooRepo, gateBar)

	mtr, err := metrics.New()
	if err != nil {
		return nil, err
	}

	lgGateIn := logrus.New().WithField("service", cnf.GateIn.Title)
	gateIn := gatein.New(startstop.New(1*time.Second), lgGateIn, cnf.GateIn, fbi, mtr)

	app := &Application{
		logger:    logrus.New().WithField("service", "app"),
		listToRun: []Item{fooRepo, gateBar, fbi, gateIn},
	}

	return app, nil
}
