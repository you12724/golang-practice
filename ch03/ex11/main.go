package main

import (
	"bytes"
	"strings"
)

func main() {
	println(comma("-123456789.123456"))
}

func comma(s string) string {
	var buf bytes.Buffer

	arr := strings.Split(s, ".")
	num := arr[0]

	if num[0] == '+' || num[0] == '-' {
		buf.WriteByte(num[0])
		num = num[1:]
	}

	for i, e := range num {
		if i != 0 && len(num)%3 == i%3 {
			buf.WriteString(",")
		}
		buf.WriteRune(e)
	}

	if len(arr) == 2 {
		buf.WriteRune('.')
		buf.WriteString(arr[1])
	}
	return buf.String()
}
