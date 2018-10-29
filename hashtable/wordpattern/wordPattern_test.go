package wordpattern

import "testing"

type Tables struct {
	pattern string
	str     string
	ok      bool
}

func TestWordPattern(t *testing.T) {
	var ok bool
	tables := []Tables{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}

	for _, v := range tables {
		if ok = wordPattern(v.pattern, v.str); ok != v.ok {
			t.Errorf("pattern=%s, str=%s, get %v, expect %v", v.pattern, v.str, ok, v.ok)
		}
	}
}
