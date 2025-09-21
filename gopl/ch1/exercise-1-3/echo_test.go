package main

import (
	"strings"
	"testing"
)

func makeArgs(n int) []string {
	args := make([]string, n)
	for i := range args {
		args[i] = "arg" + strings.Repeat("Ana Caroline <3", 10)
	}
	return args
}

func BenchmarkEchoWithFor(b *testing.B) {
	args := makeArgs(1000)

	for b.Loop() {
		_ = echoWithFor(args)
	}
}

func BenchmarkEchoWithRange(b *testing.B) {
	args := makeArgs(1000)

	for b.Loop() {
		_ = echoWithRange(args)
	}
}

func BenchmarkEcho(b *testing.B) {
	args := makeArgs(1000)

	for b.Loop() {
		_ = echo(args)
	}
}
