package main

import (
	"reflect"
	"testing"
)

func TestTogetherString1(t *testing.T) {
	input := []byte("test  test")
	input = togetherSpace(input)
	if !reflect.DeepEqual(input, []byte("test test")) {
		t.Fail()
	}
}

func TestTogetherString2(t *testing.T) {
	input := []byte("山田  太郎")
	input = togetherSpace(input)
	if !reflect.DeepEqual(input, []byte("山田 太郎")) {
		t.Fail()
	}
}
