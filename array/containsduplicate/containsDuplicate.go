// 遍历数组，如果元素不在 map 中加入，否则返回 true
package containsduplicate

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	dupMap := make(map[int]bool, 0)
	for _, v := range nums {
		if _, ok := dupMap[v]; ok == false {
			dupMap[v] = true
		} else {
			return true
		}
	}
	return false
}
