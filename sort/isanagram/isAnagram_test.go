package isanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	if isAnagram("anagram", "nagaram") != true {
		t.Errorf("Get false, Expect true")
	}

	if isAnagram("rat", "car") != false {
		t.Errorf("Get true, Expect flase")
	}

	if isAnagram("anagra\u007Am", "nagaam\u007Ar") != true {
		t.Errorf("Get false, Expect true")
	}

	if isAnagram("a", "ab") != false {
		t.Errorf("Get true, Expect flase")
	}
}
