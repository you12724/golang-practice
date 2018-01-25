package main

import "testing"

func TestExpand(t *testing.T) {

	strFunc := func(string) string {
		return "test"
	}

	for _, test := range []struct {
		input  string
		result string
	}{
		{"test$footest", "testtesttest"},
		{"$foofoofoo$foo", "testfoofootest"},
		{"    $foo $foo ", "    test test "},
	} {
		output := expand(test.input, strFunc)
		if output != test.result {
			t.Errorf("output is %s, expected %s", output, test.result)
		}
	}
}
