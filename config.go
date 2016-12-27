package main

// Microservice
// Config
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//Config structure
//When you change the structure of the `Config`, make
//sure the same changes need to be made to `config.toml`
type Config struct {
	Main    ConfMain
	Session ConfSession
}

// ConfMain - basic configuration
type ConfMain struct {
	Port string
	Name string
}

// ConfSession parameters
type ConfSession struct {
	Secure string
	Name   string
}
