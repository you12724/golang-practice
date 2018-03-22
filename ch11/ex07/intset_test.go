package intset

import "testing"

const (
	A = 100000
	B = 10000
	C = 999999
)

func BenchmarkIntSetSmallWords(b *testing.B) {
	set := IntSet{}
	for i := 0; i < b.N; i++ {
		set.Add(A)
		if !set.Has(A) {
			b.Fail()
		}
		set.Add(B)
		if !set.Has(B) {
			b.Fail()
		}
		set.Add(C)
		if !set.Has(C) {
			b.Fail()
		}
	}
}

func BenchmarkIntSetLargeWords(b *testing.B) {
	set := IntSet{make([]uint64, 100)[:]}
	for i := 0; i < b.N; i++ {
		set.Add(A)
		if !set.Has(A) {
			b.Fail()
		}
		set.Add(B)
		if !set.Has(B) {
			b.Fail()
		}
		set.Add(C)
		if !set.Has(C) {
			b.Fail()
		}
	}
}

func BenchmarkMap(b *testing.B) {
	set := make(map[int]bool)
	for i := 0; i < b.N; i++ {
		set[A] = true
		if !set[A] {
			b.Fail()
		}
		set[B] = true
		if !set[B] {
			b.Fail()
		}
		set[C] = true
		if !set[C] {
			b.Fail()
		}
	}
}
