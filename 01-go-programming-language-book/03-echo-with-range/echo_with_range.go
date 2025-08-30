package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// A new string `s` is allocated on each iteration. The GC is responsible for reclaiming the old ones.
		s += sep + arg
		sep = " "
	}

	fmt.Println("Program runned succesfully! Your args are: " + s)
}
