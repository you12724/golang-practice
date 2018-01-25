package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("words: %v\nimages: %v\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (int, int) {
	words := 0
	images := 0
	words, images = visit(words, images, n)
	return words, images
}

func visit(words, images int, n *html.Node) (int, int) {
	if n == nil {
		return words, images
	}

	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}

	if n.Type == html.TextNode {
		words += len(strings.Split(n.Data, " "))
	}
	words, images = visit(words, images, n.FirstChild)
	words, images = visit(words, images, n.NextSibling)
	return words, images
}
