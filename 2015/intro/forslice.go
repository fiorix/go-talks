// +build OMIT

package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue"}
	for n, color := range colors {
		fmt.Println(n, color)
	}
}
