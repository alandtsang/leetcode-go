package strstr

import "testing"

func TestStrStr(t *testing.T) {
	var haystack, needle string
	var result int

	haystack = "hello"
	needle = "ll"
	result = strStr(haystack, needle)
	if result != 2 {
		t.Errorf("Got %d, Expect 2", result)
	}

	haystack = "aaaaa"
	needle = "bba"
	result = strStr(haystack, needle)
	if result != -1 {
		t.Errorf("Got %d, Expect -1", result)
	}

	haystack = "aaaaa"
	needle = ""
	result = strStr(haystack, needle)
	if result != 0 {
		t.Errorf("Got %d, Expect 0", result)
	}
}
