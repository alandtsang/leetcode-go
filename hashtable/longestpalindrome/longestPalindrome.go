// 字符分为奇数和偶数情况，
// 如果得到的长度和字符串 s 相等直接返回。
// 如果长度不等，返回长度+1，偶数拼接一个奇数。
package longestpalindrome

func longestPalindrome(s string) int {
	var result int
	var sMap map[rune]int

	sMap = make(map[rune]int)
	for _, r := range s {
		sMap[r]++
		if sMap[r]%2 == 0 {
			result += 2
			sMap[r] = 0
		}
	}
	if result < len(s) {
		return result + 1
	}
	return result
}
