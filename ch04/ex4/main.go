package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5}
	rotate(test, 3)
	fmt.Printf("%v\n", test)
}

func rotate(s []int, n int) {
	if len(s) == 0 {
		return
	}

	for i := 0; i < n; i++ {
		for index := range s {
			end := index + 1
			if end == len(s) {
				continue
			}
			s[end], s[index] = s[index], s[end]
		}
	}
}
