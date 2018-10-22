// 遍历 J，将元素存储在 map 中
// 遍历 S，如果元素在 map 中，count++
package numjewelsinstones

func numJewelsInStones(J string, S string) int {
	var count int
	var jMap map[rune]bool
	jMap = make(map[rune]bool)

	for _, v := range J {
		jMap[v] = true
	}
	for _, v := range S {
		if _, ok := jMap[v]; ok {
			count++
		}
	}
	return count
}
