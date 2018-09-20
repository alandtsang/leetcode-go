package searchinsert

import "testing"

func TestSearchInsert(t *testing.T) {
	var nums []int
	var result int

	nums = []int{1, 3, 5, 6}
	result = searchInsert(nums, 5)
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	nums = []int{1, 3, 5, 6}
	result = searchInsert(nums, 2)
	if result != 1 {
		t.Errorf("Get %d, Expect 1", result)
	}

	nums = []int{1, 3, 5, 6}
	result = searchInsert(nums, 7)
	if result != 4 {
		t.Errorf("Get %d, Expect 4", result)
	}

	nums = []int{1, 3, 5, 6}
	result = searchInsert(nums, 0)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}

	nums = []int{1, 3}
	result = searchInsert(nums, 3)
	if result != 1 {
		t.Errorf("Get %d, Expect 1", result)
	}
}
