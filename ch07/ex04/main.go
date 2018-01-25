package main

import (
	"io"

	"golang.org/x/net/html"
)

type CustomReader struct {
	s string
}

func (c *CustomReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func NewReader(input string) io.Reader {
	c := CustomReader{input}
	return &c
}

func main() {
	html.Parse(NewReader("test"))
}
