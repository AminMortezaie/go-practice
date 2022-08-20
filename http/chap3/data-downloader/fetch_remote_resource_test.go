package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
The httptest.NewServer() function returns a httptest.Server object with
various fields that represent the server that was created. The only argument to
the function is an object of type http.Handler. This handler
object allows us to set up the desired handlers for the test server.
**/

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello World")
			}))
	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()
	expected := "Hello World"
	data, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(data) {
		t.Errorf("Expected response to be: %s, Got: %s", expected,
			data)
	}
}
