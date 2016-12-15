package main

// Microservice
// Config
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//Config structure
//When you change the structure of the `Config`, make
//sure the same changes need to be made to `config.toml`
//
type Config struct {
	Main   Main
	Status Status
	Error  Error
}

type Main struct {
	Port string
}

type Status struct {
	Ok       string
	NotFound string
	Timeout  string
}

type Error struct {
	FileNotFound string
}
