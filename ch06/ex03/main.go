package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
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

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word != 0 {
			count += popCount(word)
		}
	}
	return count
}

func popCount(x uint64) int {
	result := 0
	for x > 0 {
		result++
		x = x & (x - 1)
	}
	return result
}

func (s *IntSet) Remove(x int) {
	word := x / 64
	s.words[word] = 0
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	intset := IntSet{}
	words := []uint64{}
	for _, sword := range s.words {
		words = append(words, sword)
	}
	intset.words = words
	return &intset
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	tmp := s.Copy()
	tmp.IntersectWith(t)
	fmt.Printf("s: %v, tmp: %v\n", s, tmp)
	s.UnionWith(t)
	s.DifferenceWith(tmp)
}

func main() {
}
