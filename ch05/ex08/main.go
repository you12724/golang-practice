package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	outline(os.Args[1])
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, nil)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if !pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return
		}
	}
}

var depth int

func ElementByID(doc *html.Node, id string) *html.Node {
	if doc.Type == html.ElementNode {
		for _, a := range doc.Attr {
			if a.Key == "id" && a.Val == id {
				return doc
			}
		}
	}
	return nil
}

func startElement(n *html.Node) bool {
	node := ElementByID(n, os.Args[2])
	if node != nil {
		fmt.Printf("%v\n", node)
		return false
	}
	return true
}
