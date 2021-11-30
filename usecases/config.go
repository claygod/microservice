package usecases

// Microservice
// Interactor (config)
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
)

var ErrUserBadRequest = errors.New("bad request")

type Config struct {
	Title string
}

type HealthResponceStatus struct {
	fooStore int `json:"foo_store"`
	barGate  int `json:"bar_gate"`
}