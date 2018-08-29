package longestcommonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	var isExist bool
	var prefixStr string

	// find the shortest str
	index, shortestStr := findShortestStr(strs)
	if len(strs) <= 1 {
		return shortestStr
	}

	for i := 1; i <= len(shortestStr); i++ {
		for j, str := range strs {
			if j == index {
				continue
			}
			isExist = strings.HasPrefix(str, shortestStr[:i])
			if isExist == false {
				break
			}
		}

		if isExist == false {
			break
		} else {
			prefixStr = shortestStr[:i]
		}
	}
	if len(prefixStr) > 0 {
		return prefixStr
	} else {
		return ""
	}
}

func findShortestStr(strs []string) (int, string) {
	tmpStr := strs[0]
	index := 0

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(tmpStr) {
			index = i
			tmpStr = strs[i]
		}
	}
	return index, strs[index]
}
