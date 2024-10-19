package foo

// Microservice
// Foo repository
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/claygod/microservice/domain"
)

const (
	secHelthBorder = 20
	emptyString    = ""
)

/*
FooRepo - mock repository
*/
type FooRepo struct {
	hasp   domain.StartStopInterface
	logger *slog.Logger
	config *Config
}

func New(ss domain.StartStopInterface, lg *slog.Logger, cnf *Config) *FooRepo {
	return &FooRepo{
		hasp:   ss,
		logger: lg,
		config: cnf,
	}
}

func (f *FooRepo) GetFoo(fooID string) (*domain.Foo, error) {
	if !f.hasp.Add() {
		return nil, errors.New("service is stopped")
	}
	defer f.hasp.Done()

	switch {
	case fooID == "one":
		return &domain.Foo{BarID: "two"}, nil

	case fooID == "secret": // not found
		return nil, nil

	case fooID == emptyString:
		f.logger.Warn(fmt.Sprintf("%s:length of id `%s` is zero", f.config.Title, fooID))

		return nil, nil

	case len(fooID) > f.config.MaxIDLenght:
		f.logger.Warn(fmt.Sprintf("%s:length of id `%s` is greater than %d", f.config.Title, fooID, f.config.MaxIDLenght))

		return nil, nil

	default:
		return &domain.Foo{BarID: f.config.Prefix + fooID}, nil
	}
}

func (f *FooRepo) Start() error {
	if !f.hasp.Start() {
		return fmt.Errorf("%s:failed to start", f.config.Title)
	}

	return nil
}

func (f *FooRepo) Stop() error {
	if !f.hasp.Stop() {
		return fmt.Errorf("%s:failed to stop", f.config.Title)
	}

	return nil
}

func (f *FooRepo) CheckStatus() int {
	if time.Now().Second() < secHelthBorder {
		return http.StatusServiceUnavailable
	}

	return http.StatusOK
}

type Config struct {
	Title       string `toml:"REPO_TITLE" yaml:"REPO_TITLE" env:"REPO_TITLE"`
	Prefix      string `toml:"REPO_PREFIX" yaml:"REPO_PREFIX" env:"REPO_PREFIX"`
	MaxIDLenght int    `toml:"REPO_MAX_ID_LENGTH" yaml:"REPO_MAX_ID_LENGTH" env:"REPO_MAX_ID_LENGTH"`
}
