package main

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanLines)
	var count int
	for sc.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	sc := bufio.NewScanner(bytes.NewReader(p))
	sc.Split(bufio.ScanWords)
	var count int
	for sc.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

func main() {
}
