package tools

// Microservice
// Session
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// NewSession - create a new Session
func NewSession(secure string, name string) *Session {
	s := &Session{
		store: sessions.NewCookieStore([]byte(secure)),
	}
	return s
}

// Session structure
type Session struct {
	name  string
	store *sessions.CookieStore
}

// Check the run session
func (s *Session) Check(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request) {
	/*
		// Example
		session, err := s.store.Get(req, s.name)
		if err != nil {
			return w, req, true
		} else {
			session.Values["bar"] = "def"
			session.Save(req, w)
		}
	*/
	return w, req
}
