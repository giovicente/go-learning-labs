// Try comparing the execution time of each version of the echo program
// Using the time package and tests for a systematic performance evaluation

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const reps = 10000

func main() {
	args := os.Args[1:]

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

func measure(label string, fn func()) {
	start := time.Now()
	fn()
	secs := time.Since(start).Seconds()

	fmt.Printf("%s %.2fs\n", label, secs)
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
