package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"servers_api/initializers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	expectedResponse := `{"message":"pong"}`

	router := initializers.SetUpRouter()
	router.GET("/health", HealthCheckHandler)
	request, _ := http.NewRequest("GET", "/health", nil)
	httpRecorder := httptest.NewRecorder()

	router.ServeHTTP(httpRecorder, request)
	responseData, _ := io.ReadAll(httpRecorder.Body)

	assert.Equal(t, expectedResponse, string(responseData))
	assert.Equal(t, http.StatusOK, httpRecorder.Code)
}
