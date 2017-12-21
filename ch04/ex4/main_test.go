package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	test := []int{1, 2, 3, 4, 5}
	rotate(test, 3)
	if !reflect.DeepEqual(test, []int{4, 5, 1, 2, 3}) {
		t.Fail()
	}
}
