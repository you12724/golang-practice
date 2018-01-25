package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	findText()
}

func findText() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findText: %v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode && n.Data != "style" && n.Data != "script" {
		fmt.Printf("%v\n", n.Data)
	}
	texts = visit(texts, n.FirstChild)
	texts = visit(texts, n.NextSibling)
	return texts
}
