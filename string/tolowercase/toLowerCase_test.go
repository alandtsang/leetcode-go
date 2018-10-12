package tolowercase

import (
	"strings"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	var result string

	result = toLowerCase("Hello")
	if strings.Compare(result, "hello") != 0 {
		t.Errorf("Get %s, Expect hello", result)
	}
}
