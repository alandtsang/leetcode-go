package lengthoflongestsubstring

import "testing"

type Tables struct {
	s      string
	length int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tables := []Tables{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"au", 2},
	}

	var result int
	for _, v := range tables {
		if result = lengthOfLongestSubstring(v.s); result != v.length {
			t.Errorf("%s, get %d, expect %d", v.s, result, v.length)
		}
	}
}
