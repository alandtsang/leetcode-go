package isisomorphic

import "testing"

type Tables struct {
	s      string
	t      string
	expect bool
}

func TestIsIsomorphic(t *testing.T) {
	tables := []Tables{
		//{"egg", "add", true},
		//{"foo", "bar", false},
		//{"paper", "title", true},
		//{"", "", true},
		//{"aba", "baa", false},
		{"ab", "aa", false},
	}

	var result bool
	for _, v := range tables {
		if result = isIsomorphic(v.s, v.t); result != v.expect {
			t.Errorf("s=%s, t=%s, get %v, expect %v", v.s, v.t, result, v.expect)
		}
	}
}
