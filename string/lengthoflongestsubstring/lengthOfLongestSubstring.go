// 用 map 存储字符串中<字符, 位置>
// 用两个指针表示最大子串长度的首尾
// 移动右指针，并更新 map 中字符位置
// 比较移动指针后的子串长度是否大于已保存的长度
package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var maxLen int
	sMap := make(map[byte]int)
	for i, j := 0, -1; i < len(s); i++ {
		if _, ok := sMap[s[i]]; ok {
			j = max(j, sMap[s[i]])
		}
		sMap[s[i]] = i
		//fmt.Println(i, j, "key=", s[i], "value=", sMap[s[i]])
		maxLen = max(maxLen, i-j)
	}
	return maxLen
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
