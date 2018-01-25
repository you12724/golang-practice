package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	list := ElementsByTagName(doc, "div")

	for _, node := range list {
		fmt.Printf("%v\n", node)
	}
}

func ElementsByTagName(doc *html.Node, nameList ...string) []*html.Node {
	var result []*html.Node
	result = visit(result, doc, nameList...)
	return result
}

func visit(nodes []*html.Node, n *html.Node, nameList ...string) []*html.Node {
	if n == nil {
		return nodes
	}
	for _, name := range nameList {

		if n.Type == html.ElementNode && n.Data == name {
			nodes = append(nodes, n)
		}
	}
	nodes = visit(nodes, n.FirstChild, nameList...)
	nodes = visit(nodes, n.NextSibling, nameList...)
	return nodes
}
