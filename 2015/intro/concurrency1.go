// +build OMIT

package main

import "log"

// START OMIT
func main() {
	c := make(chan string)
	go writeString(c) // HL
	v := <-c          // HL
	log.Println(v)
}

func writeString(c chan string) {
	c <- "hello world" // HL
}

// STOP OMIT
