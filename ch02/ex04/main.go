package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func MyPopCount(x uint64) int {
	var result uint64
	for i := 0; i < 64; i++ {
		result += (x >> uint(i)) & 1
	}
	return int(result)
}

func main() {
	var input uint64 = 1110101041024012
	start1 := time.Now()
	fmt.Println(PopCount(input))
	fmt.Println(time.Since(start1))
	start2 := time.Now()
	fmt.Println(MyPopCount(input))
	fmt.Println(time.Since(start2))
}
