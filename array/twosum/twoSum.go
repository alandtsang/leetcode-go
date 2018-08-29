package twosum

func twoSum(nums []int, target int) []int {
	var result []int
	mapNums := make(map[int]int)

	for i, num := range nums {
		mapNums[num] = i
	}

	for i, num := range nums {
		other := target - num
		index, ok := mapNums[other]
		if ok == true && i != index {
			result = append(result, i, index)
			break
		}
	}
	return result
}
