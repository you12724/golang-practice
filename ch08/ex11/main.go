package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var done = make(chan struct{})

func main() {
	count := len(os.Args[1:])
	req := make(chan string, count)
	for _, url := range os.Args[1:] {
		go fetch(url, req)
	}
	resp := <-req
	close(done)
	fmt.Printf("\nFirst response is %s\n", resp)
}

func fetch(url string, req chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
	select {
	case <-done:
		return
	case req <- url:
	}
}
