package popcount

import "testing"

const x = 1110101041024012

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCount24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount24(x)
	}
}

func BenchmarkPopCount25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount25(x)
	}
}
