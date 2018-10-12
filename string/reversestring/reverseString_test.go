package reversestring

import (
	"strings"
	"testing"
)

func TestReverseString(t *testing.T) {
	var result string

	result = reverseString("hello")
	if strings.Compare(result, "olleh") != 0 {
		t.Errorf("Get %s, Expect olleh", result)
	}

	result = reverseString("A man, a plan, a canal: Panama")
	if strings.Compare(result, "amanaP :lanac a ,nalp a ,nam A") != 0 {
		t.Errorf("Get %s, Expect amanaP :lanac a ,nalp a ,nam A", result)
	}
}
