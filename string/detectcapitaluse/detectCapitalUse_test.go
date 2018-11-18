package detectcapitaluse

import "testing"

type Tables struct {
	word  string
	equal bool
}

func TestDetectCapitalUse(t *testing.T) {
	tables := []Tables{
		{"USA", true},
		{"leetcode", true},
		{"Google", true},
		{"FlaG", false},
		{"", false},
	}

	var ret bool
	for _, v := range tables {
		if ret = detectCapitalUse(v.word); ret != v.equal {
			t.Errorf("%q, get %v, expect %v", v.word, ret, v.equal)
		}
	}
}
