package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Idiomatic and memory efficient.
	fmt.Println(strings.Join(os.Args[1:], " "))
}
