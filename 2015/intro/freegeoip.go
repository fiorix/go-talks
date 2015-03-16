// +build OMIT

package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func main() {
	g, err := query("https://freegeoip.net/json/")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", g)
}

// START OMIT
type Geolocation struct {
	Country string `json:"country_name"`
	City    string `json:"city"`
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
