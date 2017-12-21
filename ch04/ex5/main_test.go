package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestUniq(t *testing.T) {
	s := "aaaiiiueo"
	out := uniq(strings.Split(s, ""))
	if !reflect.DeepEqual(out, []string{"a", "i", "u", "e", "o"}) {
		t.Fail()
	}
}
