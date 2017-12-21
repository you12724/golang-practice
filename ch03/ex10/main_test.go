package main

import "testing"

func TestComma(t *testing.T) {
	input := "123456789"
	output := comma(input)
	if output != "123,456,789" {
		t.Fatal()
	}
}
