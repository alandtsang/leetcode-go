package longestcommonprefix

import (
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var strs []string
	var result string

	strs = []string{"flower", "flow", "flight"}
	result = longestCommonPrefix(strs)
	if strings.Compare(result, string("fl")) != 0 {
		t.Errorf("Got %s, expect fl", result)
	}

	strs = []string{"dog", "racecar", "car"}
	result = longestCommonPrefix(strs)
	if len(result) > 0 {
		t.Errorf("Got %s, expect nil", result)
	}
}
