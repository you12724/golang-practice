package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestStartElement(t *testing.T) {
	r := strings.NewReader("<div>test</div>")
	fmt.Println(r)
	input1, _ := html.Parse(r)
	fmt.Println(input1)

	for _, test := range []struct {
		input  *html.Node
		result string
	}{
		{input1, "<div/>"},
	} {
		output := startElement(test.input)
		if test.result != output {
			t.Errorf("output is %s, expected %s", output, test.result)
		}
	}
}
