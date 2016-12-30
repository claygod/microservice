package tools

// Microservice tools
// Test Session
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"net/http"
	"testing"
)

func TestSession(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	s := NewSession("ab", "cd")
	_, req = s.Check(nil, req)
	if req == nil {
		t.Error("Error in check session")
	}
}

func TestSessionWithNil(t *testing.T) {
	s := NewSession("ab", "cd")
	_, req := s.Check(nil, nil)
	if req != nil {
		t.Error("Error in check session (with nil)")
	}
}
