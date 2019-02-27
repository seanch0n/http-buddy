package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/myip", nil)

	// In case there is an error in forming the request,
	// we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function.
	// "handler" is the handler function defined in our
	// main.go file that we want to test
	hf := http.HandlerFunc(getMyIP)

	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// Right now the test request comes in with no values. Not really
	// sure how to fix that issue
	expected := `{"IP":"","Port":""}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: "+
			"got %v want %v", actual, expected)
	}

}
