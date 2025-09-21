package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

/**
 * Test Input: gopl.io
 * Description:
 *   - A domain name without a URL scheme (http/https).
 *   - The code below will automatically fix the URL by adding the "http://" prefix if it is missing,
 *     so the program will actually fetch "http://gopl.io".
 *   - This ensures compatibility with http.Get, which requires a scheme.
 */

/**
 * Test Input: http://gopl.io
 * Description:
 *   - A complete URL with the "http" scheme.
 *   - The program should successfully fetch and print the HTTP response body from gopl.io.
 */

/**
 * Test Input: https://brasilapi.com.br/api/feriados/v1/2024
 * Description:
 *   - A full HTTPS URL pointing to an API endpoint that returns holiday data for 2024 in Brazil.
 *   - The program should fetch and print the JSON response from the API.
 */
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) && !strings.HasPrefix(url, httpsPrefix) {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("HTTP Status: %s\n", resp.Status)

		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
