package main

import "testing"

func TestComma1(t *testing.T) {
	input := "123456789"
	output := comma(input)
	if output != "123,456,789" {
		t.Fatalf("output is %s", output)
	}
}

func TestComma2(t *testing.T) {
	input := "123456789.012345"
	output := comma(input)
	if output != "123,456,789.012345" {
		t.Fatalf("output is %s", output)
	}
}

func TestComma3(t *testing.T) {
	input := "-123456789.012345"
	output := comma(input)
	if output != "-123,456,789.012345" {
		t.Fatalf("output is %s", output)
	}
}
