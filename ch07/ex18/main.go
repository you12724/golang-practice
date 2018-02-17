package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, Element{tok.Name, tok.Attr, []Node{}})
		case xml.EndElement:
			if len(stack) == 1 {
				break
			}
			stack[len(stack)-2].Children = append(stack[len(stack)-2].Children, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		case xml.CharData:
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, tok)
		}
	}
	printElement([]Node{stack[0]})
}

func printElement(nodes []Node) {
	for _, node := range nodes {
		switch node := node.(type) {
		case Element:
			println(node.Type.Local)
			printElement(node.Children)
		case CharData:
			println(node)
		}
	}
}
