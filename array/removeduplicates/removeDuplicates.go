package removeduplicates

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var i, j int = 0, 0

	for i < len(nums)-1 {
		// find the same element and jump over
		for j = i + 1; j < len(nums) && nums[i] == nums[j]; j++ {
		}
		nums = append(nums[:i+1], nums[j:]...)
		i++
		//fmt.Println(nums)
	}
	return len(nums)
}
