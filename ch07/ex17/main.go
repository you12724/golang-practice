package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []map[string]bool
	var elements []string
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
			arr := map[string]bool{}
			arr[tok.Name.Local] = true
			for _, attr := range tok.Attr {
				arr[attr.Value] = true
			}
			stack = append(stack, arr) // push
			elements = append(elements, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
			elements = elements[:len(elements)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(elements, " "), tok)
			}
		}
	}
}

func containsAll(x []map[string]bool, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0][y[0]] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
