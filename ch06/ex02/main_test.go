package main

import "testing"

func TestAddAll(t *testing.T) {
	set := new(IntSet)
	inputs := []int{1, 2, 3, 4, 5}
	set.AddAll(inputs...)
	for _, input := range inputs {
		if !set.Has(input) {
			t.Fail()
		}
	}
}
