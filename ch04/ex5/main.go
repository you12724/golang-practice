package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaaiiiueo"
	fmt.Printf("%v\n", uniq(strings.Split(s, "")))
}

func uniq(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
