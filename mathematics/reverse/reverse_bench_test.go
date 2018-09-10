package reverse

import "testing"

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(123)
		reverse(-123)
		reverse(120)
		reverse(0)
		reverse(2147483647)
		reverse(-2147483647)
	}
}

func BenchmarkReverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse2(123)
		reverse2(-123)
		reverse2(120)
		reverse2(0)
		reverse2(2147483647)
		reverse2(-2147483647)
	}
}
