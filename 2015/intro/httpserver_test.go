// +build OMIT

package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	srv := httptest.NewServer(mux)
	defer srv.Close()
	body, err := query(srv.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
	if body != "hello world\n" {
		t.Fatalf("Unexpected response: %q", body)
	}
}

// START OMIT
func query(url string) (body string, err error) { // HL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() // HL
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// STOP OMIT
