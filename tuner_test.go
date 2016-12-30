package main

// Microservice
// Test Tuner
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"reflect"
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

func TestTunerParseCommandLine(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	_, err = conf.parseCommandLine()

	if err != nil {
		t.Error("Error parsing command line with no arguments")
	}
}

func TestTunerEnvEmpty(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	m := make(map[string]string)
	err = conf.env(m)

	if err != nil {
		t.Error(err)
	}
}

func TestTunerEnvFilled(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	m := map[string]string{"Main/Port": "81"}
	err = conf.env(m)

	if conf.Main.Port != "81" {
		t.Error(err)
	}
}

func TestTunerReflecting(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	err = conf.reflecting("Main", "Port", "81")
	if err != nil {
		t.Error(err)
	}

	if conf.Main.Port != "81" {
		t.Error("Create in reflecting (error assigning a new value)")
	}
}

func TestTunerSwitchTypeInt(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	var x int = 80
	v := reflect.ValueOf(&x)
	if err := conf.switchType(v.Elem(), "81"); err != nil {
		t.Error(err)
	}
}

func TestTunerSwitchTypeFloat(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	var x float64 = 1.1
	v := reflect.ValueOf(&x)
	if err := conf.switchType(v.Elem(), "2.2"); err != nil {
		t.Error(err)
	}
}

func TestTunerSwitchTypeString(t *testing.T) {
	conf, err := NewTuner("./config.toml")
	if err != nil {
		t.Error(err)
	}
	var x string = "ab"
	v := reflect.ValueOf(&x)
	if err := conf.switchType(v.Elem(), "cd"); err != nil {
		t.Error(err)
	}
}
