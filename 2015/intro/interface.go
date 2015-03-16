// +build OMIT

package main

import "fmt"

// START OMIT
type Person interface {
	Name() string
}

func sayHi(p Person) {
	fmt.Println("Hi, I'm", p.Name())
}

type clown string

func (c clown) Name() string { return string(c) } // HL

func main() {
	sayHi(clown("Bozo"))
	sayHi(clown("Krusty"))
}

// STOP OMIT
