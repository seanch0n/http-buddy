package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// boilerplate to send http requests through testing framework
func buildReq(method string, endpoint string) (*http.Request,
	*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	recorder := httptest.NewRecorder()
	return req, recorder, nil
}

func TestGetMyIPBothIPAndPort(t *testing.T) {
	req, recorder, err := buildReq("GET", "/myip")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getMyIP)
	req.RemoteAddr = "10.10.10.10:4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getMyIP returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// It should just return what we set in @req.RemoteAddr
	expected := `{"IP":"10.10.10.10","Port":"4444"}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getMyIP returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetMyIPJustIP(t *testing.T) {
	req, recorder, err := buildReq("GET", "/myip")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getMyIP)
	req.RemoteAddr = "10.10.10.10"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getMyIP returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// Since it's just getting an IP but not a source port for whatever reason,
	// it should still return the IP rather than failing.
	expected := `{"IP":"10.10.10.10","Port":""}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getMyIP returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetMyIPJustPort(t *testing.T) {
	req, recorder, err := buildReq("GET", "/myip")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getMyIP)
	req.RemoteAddr = ":4444"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getMyIP returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// I can't image this ever happening, but if it does and we don't handle
	// it then we would crash. So probably an okay test
	expected := `{"IP":"","Port":"4444"}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getMyIP returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetMyIPNoIPOrPort(t *testing.T) {
	req, recorder, err := buildReq("GET", "/myip")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getMyIP)
	req.RemoteAddr = ""
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getMyIP returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// I can't image this ever happening, but if it does and we don't handle
	// it then we would crash. So probably an okay test
	expected := `{"IP":"","Port":""}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getMyIP returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetMyIPTooManyFields(t *testing.T) {
	req, recorder, err := buildReq("GET", "/myip")
	if err != nil {
		t.Fatal(err)
	}
	// hf is the handler to serve our request to getMyIP
	// we set our request IP/Port to a value, b/c the default is blank
	handler := http.HandlerFunc(getMyIP)
	req.RemoteAddr = "1.1.1.1:4444:3333"
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("getMyIP returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestGetHeaders(t *testing.T) {
	req, recorder, err := buildReq("GET", "/headers")
	if err != nil {
		t.Fatal(err)
	}

	// we set one header for our request
	handler := http.HandlerFunc(getHeaders)
	req.Header.Add("one", "two")
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getHeaders returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// we should get the header we sent and a count of 1
	expected := `{"Headers":[{"One":"two"}],"count":1}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getHeaders returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

// this is a bad test because the headers can be returned in any order
// which breaks this test even though it technically passes
func TestGetHeadersThree(t *testing.T) {
	req, recorder, err := buildReq("GET", "/headers")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(getHeaders)
	req.Header.Add("one", "two")
	req.Header.Add("a", "b")
	req.Header.Add("apple", "banana")
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getHeaders returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Headers":[{"One":"two"},{"A":"b"},{"Apple":"banana"}],"count":3}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getHeadersThree returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetHeadersNone(t *testing.T) {
	req, recorder, err := buildReq("GET", "/headers")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(getHeaders)
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getHeaders returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Headers":null,"count":0}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getHeadersThree returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}

func TestGetUserAgent(t *testing.T) {
	req, recorder, err := buildReq("GET", "/useragent")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(getUserAgent)
	req.Header.Add("User-Agent", "my cool user agent")
	handler.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("getHeaders returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"User-Agent":"my cool user agent"}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("getHeadersThree returned unexpected body: "+
			"got %v want %v", actual, expected)
	}
}
