// main_test

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// this acts as the recoder that we send our request to...
	recorder := httptest.NewRecorder()

	// create handler func
	hf := http.HandlerFunc(handler)

	//serve the request
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Wrong code returned: %v, expected %v", status, http.StatusOK)

		expected := "Hello World!"
		actual := recorder.Body.String()

		if actual != expected {
			t.Errorf("Handler returned unexpected : %v, expected %v", actual, expected)

		}

	}

}

func TestRouter(t *testing.T) {

	r := newRouter()

	// creating a new server
	mockServer := httptest.NewServer(r)

	// mock the request
	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Wrong response status : %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	//read body

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	//convert bytes to string
	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("wrong response : %v, expected : %v", respString, expected)
	}

}

func TestRouterFailForWrongType(t *testing.T) {

	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("wrong type of status code : %v, expected: %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("wrong response: %v, expected : %v", respString, expected)
	}
}

func TestGetCinemaByPostcode(t *testing.T) {

}
