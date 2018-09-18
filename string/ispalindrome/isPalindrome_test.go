package ispalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	var str string
	var result bool

	str = "A man, a plan, a canal: Panama"
	result = isPalindrome(str)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	str = "race a car"
	result = isPalindrome(str)
	if result != false {
		t.Errorf("Got %v, Expect false", result)
	}

	str = "0P"
	result = isPalindrome(str)
	if result != false {
		t.Errorf("Got %v, Expect false", result)
	}

	str = ""
	result = isPalindrome(str)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	str = " "
	result = isPalindrome(str)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	str = "  "
	result = isPalindrome(str)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	str = "a."
	result = isPalindrome(str)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	str = "ab2a"
	result = isPalindrome(str)
	if result != false {
		t.Errorf("Got %v, Expect false", result)
	}

	str = " apG0i4maAs::sA0m4i0Gp0"
	result = isPalindrome(str)
	if result != false {
		t.Errorf("Got %v, Expect false", result)
	}
}
