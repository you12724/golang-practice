package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	countTag()
}

func countTag() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countTag: %v\n", err)
		os.Exit(1)
	}
	counter := map[string]int{}
	for key, value := range visit(counter, doc) {
		fmt.Printf("%v: %v\n", key, value)
	}
}

func visit(counter map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return counter
	}

	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	counter = visit(counter, n.FirstChild)
	counter = visit(counter, n.NextSibling)
	return counter
}
