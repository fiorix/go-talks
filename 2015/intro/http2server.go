// +build OMIT

package main

// START OMIT
import (
	"fmt"
	"net/http"

	"github.com/bradfitz/http2"
)

func main() {
	srv := &http.Server{Addr: ":8080"}
	http2.ConfigureServer(srv, nil)
	http.HandleFunc("/hello", helloHandler)
	srv.ListenAndServeTLS("2015/intro/cert.pem", "2015/intro/key.pem")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s from %s\n", r.Proto, r.URL.Path, r.RemoteAddr)
	fmt.Fprintln(w, "hello world")
}

// STOP OMIT
