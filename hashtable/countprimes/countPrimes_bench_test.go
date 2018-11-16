// go test -v -bench=. -run=none
package countprimes

import "testing"

func BenchmarkCountPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPrimes(30)
	}
}

func BenchmarkCountPrimes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPrimes2(30)
	}
}

func BenchmarkCountPrimes3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPrimes3(30)
	}
}
