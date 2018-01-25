package main

import "fmt"

func main() {
	set := new(IntSet)
	set.AddAll(1, 2, 3, 10, 11, 100)
	fmt.Printf("%v", set.Elems())
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(inputs ...int) {
	for _, input := range inputs {
		s.Add(input)
	}
}

func (s *IntSet) Elems() []int {
	result := []int{}
	for i, word := range s.words {
		word := word
		count := 0
		for word > 0 {
			if word&1 == 1 {
				result = append(result, i*64+count)
			}
			word = word >> 1
			count++
		}
	}
	return result
}
