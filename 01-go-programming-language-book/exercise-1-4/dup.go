// Modify the program so that it displays the name of the files
// in which each duplicated line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "(stdin)", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			_ = f.Close()
		}
	}
	for line, perFile := range counts {
		totalCount := 0
		for _, c := range perFile {
			totalCount += c
		}
		if totalCount > 1 {
			fmt.Printf("%d\t%s\t", totalCount, line)
			for fname := range perFile {
				fmt.Printf("%s\t", fname)
			}
		}
		fmt.Println()
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Read error (%s): %v\n", filename, err)
	}
}
