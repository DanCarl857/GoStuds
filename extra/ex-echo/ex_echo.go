// Exercises to the echo program
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""

	// Exercise 1.1
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])

	// Exercise 1.2
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
