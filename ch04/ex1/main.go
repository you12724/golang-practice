package main

import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	var count int

	for i, e := range c1 {
		if e != c2[i] {
			count++
		}
	}

	println(count)
}
