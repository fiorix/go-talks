// +build OMIT

package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// const ServerURL = "https://freegeoip.net/json/"

func main() {
	g, err := query(ServerURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", g)
}

// START OMIT
const ServerURL = "http://localhost:9999/json/github.com"

type Geolocation struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

func query(url string) (*Geolocation, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	g := new(Geolocation)
	if err := json.NewDecoder(resp.Body).Decode(g); err != nil { // HL
		return nil, err
	}
	return g, nil
}

// STOP OMIT
