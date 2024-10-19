package usecases

// Microservice
// Interactor (config)
// Copyright © 2021-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
)

var ErrUserBadRequest = errors.New("bad request")

type Config struct {
	Title string `toml:"INTER_TITLE" yaml:"INTER_TITLE" env:"INTER_TITLE"`
}

type HealthResponceStatus struct {
	FooStore int `json:"foo_store"`
	BarGate  int `json:"bar_gate"`
}
