package split

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, test := range []struct {
		input     string
		separator string
		result    []string
	}{
		{"a,i,u,e,o", ",", []string{"a", "i", "u", "e", "o"}},
	} {
		output := strings.Split(test.input, test.separator)
		if !reflect.DeepEqual(output, test.result) {
			t.Errorf("Split(%v, %v) returned %v, want %v", test.input, test.separator, output, test.result)
		}
	}
}
