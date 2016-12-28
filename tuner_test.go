package main

// Microservice
// Test Tuner
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"testing"
)

func TestTunerNormArg(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}

	if conf.Main.Name != "Microservice" {
		t.Error("Create a configuration error")
	}
}

func TestTunerWrongArg(t *testing.T) {
	_, err := NewTuner("./abc.toml")
	if err == nil {
		t.Error("Error Not received when sending a wrong argument")
	}
}
