package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that a GET request to the health endpoint
func TestHealthEndpoint(t *testing.T) {
	r := getRouter()
	hc := new(HealthController)

	// Define the route similar to its definition in the routes file
	r.GET("/v1/health", hc.Retrieve)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/v1/health", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the status is "ok"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "ok") > 0

		return statusOK && pageOK
	})
}
