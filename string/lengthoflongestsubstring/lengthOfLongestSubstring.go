// 取 s 第 1 个字符，使用双指针向后取字符，直到字符重复。
// 继续 s 的下个字符，重复上一动作。
package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var length = 1
	for i := 0; i < len(s); i++ {
		sMap := make(map[byte]bool)
		sMap[s[i]] = true
		var tempLen = 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := sMap[s[j]]; ok {
				break
			} else {
				sMap[s[j]] = true
				tempLen++
			}
		}
		if tempLen > length {
			length = tempLen
		}
	}
	return length
}
