package main

import "testing"

func TestJoin(t *testing.T) {
	for _, test := range []struct {
		input  []string
		sep    string
		result string
	}{
		{[]string{""}, ",", ""},
		{[]string{"1", "2"}, ",", "1,2"},
		{[]string{"111", "124", "あいう"}, "-", "111-124-あいう"},
	} {
		output := join(test.sep, test.input...)
		if test.result != output {
			t.Errorf("output is %s, expected %s", output, test.result)
		}
	}
}
