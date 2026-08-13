package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPUT(t *testing.T)  {
	body := strings.NewReader(`{"value": "tolu"}`)
	req := httptest.NewRequest("PUT", "/key/username", body)
	rr := httptest.NewRecorder()

	store := &store{data: make(map[string]string)}
	store.PUT(rr, req)

	if rr.Code != 200 {
		t.Errorf("Expected status code %v, got %v", 200, rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "stored") {
		t.Errorf("Expected {status: stored}, got %v", rr.Body.String())
	}
}