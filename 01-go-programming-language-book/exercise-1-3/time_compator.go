// Try comparing the execution time of each version of the echo program
// Using the time package and tests for a systematic performance evaluation

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const reps = 10000000

func main() {
	args := os.Args[1:]

	printTitle()

	measure("FOR ->", func() {
		var sink string
		for i := 0; i < reps; i++ {
			sink = echoWithFor(args)
		}
		_ = sink
	})

	measure("RANGE ->", func() {
		var sink string
		for i := 0; i < reps; i++ {
			sink = echoWithRange(args)
		}
		_ = sink
	})

	measure("IDIOMATIC ->", func() {
		var sink string
		for i := 0; i < reps; i++ {
			sink = echo(args)
		}
		_ = sink
	})
}

func printTitle() {
	fmt.Printf("=== LOOP TIME BENCHMARK IN MICROSSECONDS ===\n")
}

func measure(label string, fn func()) {
	start := time.Now()
	fn()
	ms := time.Since(start).Microseconds()

	fmt.Printf("%s %dμs\n", label, ms)
}

func echoWithFor(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	return s
}

func echoWithRange(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	return s
}

func echo(args []string) string {
	return strings.Join(args, " ")
}
