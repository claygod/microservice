package usecases

// Microservice
// Interactor (config)
// Copyright © 2021-2024 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
)

var ErrUserBadRequest = errors.New("bad request")

type Config struct {
	Title *string `env:"USECASES_TITLE"`
}

type HealthResponceStatus struct {
	fooStore int `json:"foo_store"`
	barGate  int `json:"bar_gate"`
}
