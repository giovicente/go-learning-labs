package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	var results []string

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a gorroutine
	}

	for range os.Args[1:] {
		result := <-ch
		results = append(results, result)
	}

	output := ""
	for _, r := range results {
		output += r + "\n"
		fmt.Println(r)
	}

	err := os.WriteFile("out.txt", []byte(output), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	var buf bytes.Buffer

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	var nbytes int64
	nbytes, err = io.Copy(&buf, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
