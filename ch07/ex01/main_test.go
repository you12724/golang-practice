package main

import "testing"

func TestLineCounter(t *testing.T) {
	for _, test := range []struct {
		input  []byte
		result int
	}{
		{[]byte("hello, bob"), 1},
		{[]byte("hello\nbob"), 2},
		{[]byte("\n\n\n"), 3},
		{[]byte(""), 0},
	} {
		var c LineCounter
		output, _ := c.Write(test.input)
		if output != test.result {
			t.Errorf("output is %d, expected %d", output, test.result)
		}
	}
}

func TestWordCounter(t *testing.T) {
	for _, test := range []struct {
		input  []byte
		result int
	}{
		{[]byte("hello, bob"), 2},
		{[]byte("hello\nbob"), 2},
		{[]byte("nice to meet you"), 4},
		{[]byte(""), 0},
	} {
		var c WordCounter
		output, _ := c.Write(test.input)
		if output != test.result {
			t.Errorf("output is %d, expected %d", output, test.result)
		}
	}
}
