package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code %v harusnya %v", status, http.StatusOK)
	}

	expected := `{"message":"service available!"}`
	if rr.Body.String() != expected {
		t.Errorf("body: %v harusnya %v",
			rr.Body.String(), expected)
	}
}
