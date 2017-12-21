package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	input := [5]int{1, 2, 3, 4, 5}
	reverse(&input)
	if !reflect.DeepEqual(input, [5]int{5, 4, 3, 2, 1}) {
		t.Fail()
	}
}
