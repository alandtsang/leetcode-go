package countprimes

func countPrimes(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// Check if it is not divisible by any number less than n.
// complexity O(n)
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes2(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isPrime2(i) {
			count++
		}
	}
	return count
}

// The number must not be divisible by any number > n / 2
func isPrime2(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes3(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isPrime3(i) {
			count++
		}
	}
	return count
}

func isPrime3(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
