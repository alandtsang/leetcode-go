package firstuniqchar

import "testing"

type Tables struct {
	s   string
	idx int
}

func TestFirstUniqChar(t *testing.T) {
	var result int
	tables := []Tables{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}

	for _, v := range tables {
		if result = firstUniqChar(v.s); result != v.idx {
			t.Errorf("s=%s, result=%d, expect %d", v.s, result, v.idx)
		}
	}
}
