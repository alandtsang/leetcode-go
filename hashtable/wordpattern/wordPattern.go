// 利用 map 存储字符出现的位置
// 如果 map[pattern[i]] != map[str[i]] 返回 false
package wordpattern

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	s := strings.Fields(str)
	if len(pattern) != len(s) {
		return false
	}

	patternMap := make(map[byte]int)
	strMap := make(map[string]int)
	for i := 0; i < len(pattern); i++ {
		if patternMap[pattern[i]] != strMap[s[i]] {
			return false
		}
		patternMap[pattern[i]] = i + 1
		strMap[s[i]] = i + 1
	}
	return true
}
