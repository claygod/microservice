package domain

// Microservice
// Interfaces
// Copyright © 2021 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type RepoInterface interface {
	GetFoo(ID string) (*Foo, error)
	CheckStatus() int
}

type ExtenalGateInterface interface {
	GetBar(ID string) (*Bar, error)
	CheckStatus() int
}

type StartStopInterface interface {
	Start() bool
	Stop() bool
	Add() bool
	Done() bool
	Total() int64
	IsReady() bool
	IsRun() bool
}
