// 变量 s 和 t，对比包含字符的个数即可，注意两个 map 的长短。
// go 可以使用 rune 来存储
package isanagram

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return true
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		tMap[v]++
	}

	if len(sMap) > len(tMap) {
		for k, v := range sMap {
			if tMap[k] != v {
				return false
			}
		}
	} else {
		for k, v := range tMap {
			if sMap[k] != v {
				return false
			}
		}
	}
	return true
}
