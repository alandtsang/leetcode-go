// 正确的单词需要满足一下3点：
// 1. 全大写（单词转换成大写后与原单词相等）
// 2. 全小写（单词转换成小写后与原单词相等）
// 3. 首字母大写其余全小写（除首字母外其余字母转换成小写后与原单词的除首字母外字母相等）
package detectcapitaluse

import "strings"

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return false
	}
	return strings.ToUpper(word) == word ||
		strings.ToLower(word) == word ||
		strings.ToLower(word[1:]) == word[1:]
}
