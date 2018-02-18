package main

import "testing"

func BenchmarkUseGoRoutine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useGoRoutine()
	}
}

func BenchmarkNoUseGoRoutine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		notUseGoRoutine()
	}
}
