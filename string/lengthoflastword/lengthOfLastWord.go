package lengthoflastword

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	strs := strings.Fields(s)
	if len(strs) == 0 {
		return 0
	}
	return len(strs[len(strs)-1])
}
