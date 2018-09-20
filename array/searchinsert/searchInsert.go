package searchinsert

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	} else if target > nums[len(nums)-1] { // last element index in nums
		return len(nums)
	}

	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return 0
}
