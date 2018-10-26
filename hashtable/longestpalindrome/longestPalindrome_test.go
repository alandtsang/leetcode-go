package longestpalindrome

import "testing"

type Tables struct {
	s      string
	expect int
}

func TestLongestPalindrome(t *testing.T) {
	var result int
	tables := []Tables{
		{"abccccdd", 7},
		{"ccc", 3},
		{"bananas", 5},
	}

	for _, v := range tables {
		if result = longestPalindrome(v.s); result != v.expect {
			t.Errorf("%s, get %d, expect %d", v.s, result, v.expect)
		}
	}
}
