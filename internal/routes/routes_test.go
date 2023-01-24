package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	basePath := "/api/v1"
	request := httptest.NewRequest(http.MethodGet, "localhost:1234"+basePath+"/ping", nil)
	response := httptest.NewRecorder()

	ping(response, request)

	expected := "pong\n"
	responseText := response.Body.String()
	if responseText != expected {
		t.Errorf("wrong response to /ping request: expected: %s, got: %s", expected, responseText)
	}
}
