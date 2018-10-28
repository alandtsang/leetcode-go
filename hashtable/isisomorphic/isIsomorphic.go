// 用 map 记录字符出现的位置，
// 如果对应的字符的位置不同则返回 false
package isisomorphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return true
	}

	var sMap, tMap map[byte]int
	sMap = make(map[byte]int)
	tMap = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]] = i + 1
		tMap[t[i]] = i + 1
	}
	return true
}
