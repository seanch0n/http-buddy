package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetGetRespGETRequest(t *testing.T) {
	req, recorder, err := buildReq("GET", "/methods/get")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getGETResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getGETResp returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `{"GET":true}`
	actual := strings.TrimSpace(recorder.Body.String())
	if actual != expected {
		t.Errorf("getGETResp returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetGetRespNotGetRequest(t *testing.T) {
	req, recorder, err := buildReq("POST", "/methods/get")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getGETResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("getGETResp returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `BAD METHOD`
	actual := strings.TrimSpace(recorder.Body.String())
	if strings.TrimSpace(actual) != expected {
		t.Errorf("getGETResp returned unexpected body: "+
			"got '%v' want '%v'", actual, expected)
	}
}

func TestGetPostReq(t *testing.T) {
	req, recorder, err := buildReq("POST", "/methods/post")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getPOSTResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getPOSTResp returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `{"POST":true}`
	actual := strings.TrimSpace(recorder.Body.String())
	if actual != expected {
		t.Errorf("getPOSTResp returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetPostRespNotPost(t *testing.T) {
	req, recorder, err := buildReq("GET", "/methods/post")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getPOSTResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("getPOSTResp returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `BAD METHOD`
	actual := strings.TrimSpace(recorder.Body.String())
	if strings.TrimSpace(actual) != expected {
		t.Errorf("getPOSTResp returned unexpected body: "+
			"got '%v' want '%v'", actual, expected)
	}
}

func TestGetPatchReq(t *testing.T) {
	req, recorder, err := buildReq("PATCH", "/methods/patch")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getPATCHResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getPATCHResp returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `{"PATCH":true}`
	actual := strings.TrimSpace(recorder.Body.String())
	if actual != expected {
		t.Errorf("getPATCHResp returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetPatchRespNotPatch(t *testing.T) {
	req, recorder, err := buildReq("GET", "/methods/patch")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getPOSTResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("getPATCHResp returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `BAD METHOD`
	actual := strings.TrimSpace(recorder.Body.String())
	if strings.TrimSpace(actual) != expected {
		t.Errorf("getPATCHResp returned unexpected body: "+
			"got '%v' want '%v'", actual, expected)
	}
}

func TestGetPutReq(t *testing.T) {
	req, recorder, err := buildReq("PUT", "/methods/put")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getPUTResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getPUTResp returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `{"PUT":true}`
	actual := strings.TrimSpace(recorder.Body.String())
	if actual != expected {
		t.Errorf("getPUTResp returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetPutRespNotPut(t *testing.T) {
	req, recorder, err := buildReq("GET", "/methods/put")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getPUTResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("getPUTResp returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `BAD METHOD`
	actual := strings.TrimSpace(recorder.Body.String())
	if strings.TrimSpace(actual) != expected {
		t.Errorf("getPUTResp returned unexpected body: "+
			"got '%v' want '%v'", actual, expected)
	}
}

func TestGetDeleteReq(t *testing.T) {
	req, recorder, err := buildReq("DELETE", "/methods/delete")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getDELETEResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getDELETEResp returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `{"DELETE":true}`
	actual := strings.TrimSpace(recorder.Body.String())
	if actual != expected {
		t.Errorf("getDELETEResp returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetDeleteRespNotDelete(t *testing.T) {
	req, recorder, err := buildReq("GET", "/methods/delete")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getDELETEResp)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("getDELETEResp returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `BAD METHOD`
	actual := strings.TrimSpace(recorder.Body.String())
	if strings.TrimSpace(actual) != expected {
		t.Errorf("getDELETEResp returned unexpected body: "+
			"got '%v' want '%v'", actual, expected)
	}
}
