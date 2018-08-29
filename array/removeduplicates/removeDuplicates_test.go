package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var nums []int
	var length int

	nums = []int{1, 1, 2}
	length = removeDuplicates(nums)
	if length != 2 {
		t.Errorf("Get %d, Expect 2", length)
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length = removeDuplicates(nums)
	if length != 5 {
		t.Errorf("Get %d, Expect 5", length)
	}
}
