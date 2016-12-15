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

// Main - basic configuration
type Main struct {
	Port string
}

// Status - types of statuses
type Status struct {
	Ok       string
	NotFound string
	Timeout  string
}

// Error - types of errors
type Error struct {
	FileNotFound string
}
