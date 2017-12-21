package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq()
}

func wordfreq() {
	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)

	counts := make(map[string]int)

	for sc.Scan() {
		counts[sc.Text()]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

}
