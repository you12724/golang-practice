package main

import "unicode"

func main() {

}

func togetherSpace(input []byte) []byte {
	isSpace := false
	for i, e := range input {
		if unicode.IsSpace(rune(e)) {
			if isSpace == true {
				copy(input[i:], input[i+1:])
				input = input[:len(input)-1]
			}
			isSpace = true
			continue
		}
		isSpace = false
	}
	return input
}
