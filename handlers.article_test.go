package main

import (
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"net/http"
	"strings"
)

func TestShowIndexPageUnauth(t *testing.T) {
	r := getRouter(true)

	r.Handle("GET", "/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200 (StatusOK)
		// Assign statusOK to the true/false evaluation of whether the return code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries to
		// parse and process HTML
		// Assign pageOK to true/false evaluation of whether err is nil & title is Home Page
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}