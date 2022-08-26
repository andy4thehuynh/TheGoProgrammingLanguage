package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(sanitize(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		printResponse(resp, url)
	}
}

// Exercise 1.8
func sanitize(url string) (res string) {
	prefix := strings.HasPrefix(url, "http://")

	if prefix {
		res = url
	} else {
		res = "http://" + url
	}
	return res
}

func printResponse(resp *http.Response, url string) {
	// Exercise 1.7
	b, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
