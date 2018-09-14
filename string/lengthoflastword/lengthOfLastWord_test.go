package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	var str string
	var result int

	str = "Hello World"
	result = lengthOfLastWord(str)
	if result != 5 {
		t.Errorf("Got %d, Expect 5", result)
	}

	str = ""
	result = lengthOfLastWord(str)
	if result != 0 {
		t.Errorf("Got %d, Expect 0", result)
	}

	str = " "
	result = lengthOfLastWord(str)
	if result != 0 {
		t.Errorf("Got %d, Expect 1", result)
	}
}
