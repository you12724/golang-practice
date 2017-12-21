package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	input := []byte("aiueo")
	input = reverse(input)
	fmt.Printf("%v\n%v\n", input, []byte("oeuia"))
	if !reflect.DeepEqual(input, []byte("oeuia")) {
		t.Fail()
	}
}
