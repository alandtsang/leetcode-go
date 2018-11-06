// 遍历 s，计算每个字符出现的次数存入 map，
// 再次遍历 s，找出字符为 1 的索引
package firstuniqchar

func firstUniqChar(s string) int {
	var sMap map[rune]int
	sMap = make(map[rune]int)

	for _, v := range s {
		sMap[v]++
	}
	for i, v := range s {
		if sMap[v] == 1 {
			return i
		}
	}
	return -1
}
