package main

// Microservice
// Test Config
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	conf, err := NewTuner("config.toml")

	if err != nil {
		t.Error("Error in `NewTuner`")
	}

	if reflect.TypeOf(conf) != reflect.TypeOf(&Tuner{}) {
		t.Error("Error in `Config")
	}
}

func TestConfigNonExistentFile(t *testing.T) {
	_, err := NewTuner("abc.toml")

	if err == nil {
		t.Error("When a non-existent file must return an error")
	}
}

func TestConfigParam(t *testing.T) {
	conf, err := NewTuner("config.toml")
	if err != nil {
		t.Error("Error in `NewTuner`")
	}

	if conf.Main.Port != "80" {
		t.Error("Error in param`Port")
	}
}
