package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
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

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) string) {
	if pre != nil {
		str := pre(n)
		fmt.Print(str)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		fmt.Print(post(n))
	}
}

var depth int

func startElement(n *html.Node) string {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			return fmt.Sprintf("%*s<%s/>\n", depth*2, "", n.Data)
		}
		result := fmt.Sprintf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
		return result
	} else {
		return fmt.Sprintf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) string {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		return fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return ""
}
