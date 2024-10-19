package app

// Microservice
// App build
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"log/slog"
	"time"

	"github.com/claygod/tools/startstop"

	"github.com/claygod/microservice/services/gateways/bar"
	"github.com/claygod/microservice/services/gateways/gatein"
	"github.com/claygod/microservice/services/metrics"
	"github.com/claygod/microservice/services/repositories/foo"
	"github.com/claygod/microservice/usecases"
)

func New(cnf *Config) (*Application, error) {
	slog.With("service", cnf.FooRepo.Title)
	fooRepo := foo.New(startstop.New(1*time.Second), slog.With("service", cnf.FooRepo.Title), cnf.FooRepo)
	gateBar := bar.New(startstop.New(1*time.Second), slog.With("service", cnf.FooRepo.Title), cnf.GateBar)
	fbi := usecases.NewFooBarInteractor(startstop.New(1*time.Second), cnf.Interactor, fooRepo, gateBar)

	mtr, err := metrics.New()
	if err != nil {
		return nil, err
	}

	gateIn := gatein.New(startstop.New(1*time.Second), slog.With("service", cnf.GateIn.Title), cnf.GateIn, fbi, mtr)

	app := &Application{
		logger:    slog.With("service", "app"),
		listToRun: []Item{fooRepo, gateBar, fbi, gateIn},
	}

	return app, nil
}
