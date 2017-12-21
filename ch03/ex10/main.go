package main

import "bytes"

func main() {
	println(comma("123456789"))
}

func comma(s string) string {
	var buf bytes.Buffer

	for i, e := range s {
		if i != 0 && len(s)%3 == i%3 {
			buf.WriteString(",")
		}
		buf.WriteRune(e)
	}
	return buf.String()
}
