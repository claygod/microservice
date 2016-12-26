package tools

// Microservice
// Logger
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"log"
)

// NewLogger - create a new Logger
func NewLogger() *Logger {
	l := &Logger{}
	return l
}

// Logger structure
type Logger struct {
	codError map[int]string
}

// Message - new message in the log
func (l *Logger) Message() *Msg {
	em := &Msg{}
	return em
}

// Msg - message struct
type Msg struct {
	buf []interface{}
}

// Field - log context
func (e *Msg) Field(key string, value interface{}) *Msg {
	e.buf = append(e.buf, key, `="`, value, `" `)
	return e
}

// Send - sending log
func (e *Msg) Send() {
	log.Print(e.buf...)
}
