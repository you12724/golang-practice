package main

import (
	"reflect"
	"testing"
)

func TestElems(t *testing.T) {
	for _, test := range []struct {
		input []int
	}{
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 21, 30, 400, 50000}},
		{[]int{1, 2, 3000, 4000, 500000}},
		{[]int{0, 2, 30000000000, 4, 500000000000000000}},
		{[]int{1, 2, 3, 4, 5}},
	} {
		set := new(IntSet)
		set.AddAll(test.input...)

		if !reflect.DeepEqual(set.Elems(), test.input) {
			t.Errorf("output is %v, expected %v", set.Elems(), test.input)
		}
	}
}
