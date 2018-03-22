package main

import (
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	for _, test := range []struct {
		input           string
		targetRune      rune
		targetRuneCount int
		targetByte      int
		targetByteCount int
		invalidCount    int
	}{
		{"aiueokakikukeko", 'k', 5, 1, 15, 0},
		{"アイウエオカキクケコ", 'ア', 1, 3, 10, 0},
		{"アイウエオカキクケコ", 'サ', 0, 1, 0, 0},
		{"I am Yosuke Hori. I'm software Engineer.", '.', 2, 1, 40, 0},
	} {
		reader := strings.NewReader(test.input)
		counts, utflen, invalid, err := charcount(reader)
		if err != nil {
			t.Fatalf("%v\n", err)
		}

		if counts[test.targetRune] != test.targetRuneCount {
			t.Fatalf("%v's count is %v, expected %v", test.targetRune, counts[test.targetRune], test.targetRuneCount)
		}

		if utflen[test.targetByte] != test.targetByteCount {
			t.Fatalf("%vbyte count is %v, expected %v", test.targetByte, utflen[test.targetByte], test.targetByteCount)
		}

		if invalid != test.invalidCount {
			t.Fatalf("invalid count is %v, expected %v", invalid, test.invalidCount)
		}
	}
}
