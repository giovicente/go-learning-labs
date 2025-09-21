package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// The for loop below starts at index 1 because index 0 of os.Args contains the program name. In Go, as in most languages, it’s also normal for indexes to start at 0.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("Program runned succesfully! Your args are: " + s)
}
