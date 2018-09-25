package addstrings

import (
	"strings"
	"testing"
)

func TestAddStrings(t *testing.T) {
	var num1, num2, result string

	num1 = "1"
	num2 = "23"
	result = addStrings(num1, num2)
	if strings.Compare(result, "24") != 0 {
		t.Errorf("Get %s, Expect 24", result)
	}

	num1 = "1"
	num2 = "9"
	result = addStrings(num1, num2)
	if strings.Compare(result, "10") != 0 {
		t.Errorf("Get %s, Expect 10", result)
	}
}
