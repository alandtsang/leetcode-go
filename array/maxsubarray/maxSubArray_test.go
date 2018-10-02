package maxsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	var nums []int
	var result int

	nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result = maxSubArray(nums)
	if result != 6 {
		t.Errorf("Get %d, Expect 6", result)
	}
}
