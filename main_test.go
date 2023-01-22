package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreating(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1234/hi", nil)
	response := httptest.NewRecorder()

	greating(response, request)

	expected := "Hi, Bala!\n"
	responseText := response.Body.String()
	if responseText != expected {
		t.Errorf("wrong response from server: expected: %s, got: %s", expected, responseText)
	}
}

func TestParting(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:1234/bye", nil)
	response := httptest.NewRecorder()

	parting(response, request)

	expected := "Bye, Bala!\n"
	responseText := response.Body.String()
	if responseText != expected {
		t.Errorf("wrong response from server: expected: %s, got: %s", expected, responseText)
	}
}
