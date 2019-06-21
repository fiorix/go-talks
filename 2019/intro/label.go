// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int)
	go func() {
		for {
			c <- rand.Intn(10)
		}
	}()
Loop:
	for v := range c {
		switch {
		case v < 9:
			fmt.Println("Got", v)
		default:
			fmt.Println("bye!")
			break Loop
		}
	}
}
