// +build OMIT

package main

// START OMIT
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// STOP OMIT
