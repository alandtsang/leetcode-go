package findthedifference

func findTheDifference(s string, t string) byte {
	var ret byte
	var sMap map[string]int

	sMap = make(map[string]int, 0)
	for _, v := range s { // 统计 s 中字符出现的次数
		sMap[string(v)] += 1
	}

	for i, v := range t {
		_, ok := sMap[string(v)]
		if !ok || sMap[string(v)] == 0 {
			return t[i]
		} else {
			sMap[string(v)] -= 1
		}
	}
	return ret
}
