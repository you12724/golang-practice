package main

import "fmt"

func main() {
	var test []int
	_ = append(test, 1)
	fmt.Printf("%v", test)
}

func tmpreverse(input []byte) []byte {
	for i := range input {
		input[i], input[len(input)-i-1] = input[len(input)-i-1], input[i]
	}
	return input
}

func reverse(input []byte) []byte {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}
