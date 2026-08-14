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

func TestGET(t *testing.T)  {
	body := strings.NewReader(`{"value": "tolu"}`)
	req := httptest.NewRequest("PUT", "/key/username", body)
	rr := httptest.NewRecorder()

	store := &store{data: make(map[string]string)}
	store.PUT(rr, req)

	req2 := httptest.NewRequest("GET", "/key/username", body)
	rr2 := httptest.NewRecorder()

	store.GET(rr2, req2)
	if !strings.Contains(rr2.Body.String(), "tolu") {
		t.Errorf("Expected {value: tolu, got %v", rr2.Body.String())
	}

	req3 := httptest.NewRequest("GET", "/key/nonexistent", nil)
	rr3 := httptest.NewRecorder()

	store.GET(rr3, req3)
	if rr3.Code != 404 {
		t.Errorf("Expected status code 404, got %v", rr3.Code)
	}
}

func TestDELETE(t *testing.T)  {
	body := strings.NewReader(`{"value": "tolu"}`)
	req := httptest.NewRequest("PUT", "/key/username", body)
	rr := httptest.NewRecorder()

	store := &store{data: make(map[string]string)}
	store.PUT(rr, req)

	req2 := httptest.NewRequest("DELETE", "/key/username", body)
	rr2 := httptest.NewRecorder()

	store.DELETE(rr2, req2)
	if !strings.Contains(rr2.Body.String(), "deleted") {
		t.Errorf("Expected {status: deleted}, got %v", rr2.Body.String())
	}

	req3 := httptest.NewRequest("DELETE", "/key/nonexistent", nil)
	rr3 := httptest.NewRecorder()

	store.DELETE(rr3, req3)
	if rr3.Code != 404 {
		t.Errorf("Expected status code 404, got %v", rr3.Code)
	}
}