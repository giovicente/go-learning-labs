// Modify the echowithfor program to also display os.Args[0]

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("Program runned succesfully! Your args are: " + s)
}
