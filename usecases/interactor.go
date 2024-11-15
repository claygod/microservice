package usecases

// Microservice
// Interactor
// Copyright © 2021-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"context"
	"fmt"

	"github.com/claygod/microservice/domain"
)

type FooBarInteractor struct {
	hasp     domain.StartStopInterface
	config   *Config
	fooStore domain.RepoInterface
	barGate  domain.ExtenalGateInterface
}

func NewFooBarInteractor(ss domain.StartStopInterface, cnf *Config, fs domain.RepoInterface, bg domain.ExtenalGateInterface) *FooBarInteractor {
	return &FooBarInteractor{
		hasp:     ss,
		config:   cnf,
		fooStore: fs,
		barGate:  bg,
	}
}

func (f *FooBarInteractor) GetBar(ctx context.Context, key string) (*domain.Bar, error) {
	if !f.hasp.Add() {
		return nil, fmt.Errorf("%s:service is stopped", f.config.Title)
	}
	defer f.hasp.Done()

	foo, err := f.fooStore.GetFoo(key)

	if foo == nil || err != nil {
		return nil, err
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return f.barGate.GetBar(foo.BarID)
}

func (f *FooBarInteractor) Start() error {
	if !f.hasp.Start() {
		return fmt.Errorf("%s:failed to start", f.config.Title)
	}

	return nil
}

func (f *FooBarInteractor) Stop() error {
	if !f.hasp.Stop() {
		return fmt.Errorf("%s:failed to stop", f.config.Title)
	}

	return nil
}

func (f *FooBarInteractor) GetHealth() *HealthResponceStatus {
	return &HealthResponceStatus{
		FooStore: f.fooStore.CheckStatus(),
		BarGate:  f.barGate.CheckStatus(),
	}
}
