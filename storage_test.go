package main

// Microservice
// Test Storage
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"testing"
)

func TestStorage(t *testing.T) {
	conf, err := NewTuner("config.toml")
	if err != nil {
		panic(err)
	}
	store := NewStorage(conf)

	if store == nil {
		t.Error("Error while creating a storage")
	}
}
