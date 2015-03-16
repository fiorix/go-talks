// +build OMIT

package main

import (
	"log"
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().UnixNano()) }

// START OMIT
func main() {
	c := make(chan string)
	go writeString(c)
	select {
	case v := <-c: // HL
		log.Println(v)
	case <-time.After(10 * time.Millisecond): // HL
		log.Println("timed out")
	}
}

func writeString(c chan string) {
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	c <- "hello world"
}

// STOP OMIT
