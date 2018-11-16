package countprimes

import "testing"

type Tables struct {
	n   int
	ret int
}

var tables = []Tables{
	{10, 4},
	{30, 10},
}

func TestCountPrimes(t *testing.T) {
	var result int
	for _, v := range tables {
		if result = countPrimes(v.n); result != v.ret {
			t.Errorf("countPrimes(%d), get %d, expect %d", v.n, result, v.ret)
		}
	}
}

func TestCountPrimes2(t *testing.T) {
	var result int
	for _, v := range tables {
		if result = countPrimes2(v.n); result != v.ret {
			t.Errorf("countPrimes2(%d), get %d, expect %d", v.n, result, v.ret)
		}
	}
}

func TestCountPrimes3(t *testing.T) {
	var result int
	for _, v := range tables {
		if result = countPrimes3(v.n); result != v.ret {
			t.Errorf("countPrimes2(%d), get %d, expect %d", v.n, result, v.ret)
		}
	}
}
