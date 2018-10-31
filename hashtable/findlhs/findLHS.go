// 遍历数组，记录数字出现的次数到 map 中，
// 遍历 map，如果对于 k 存在 map[k+1]，计算 v+v1
// 对比 max 和 v+v1
package findlhs

func findLHS(nums []int) int {
	var max int
	var numsMap map[int]int

	numsMap = make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}

	for k, v := range numsMap {
		if v1, ok := numsMap[k+1]; ok {
			if max < v+v1 {
				max = v + v1
			}
		}
	}
	return max
}
