package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	var result int
	var nums []int

	nums = []int{3, 2, 2, 3}
	result = removeElement(nums, 3)
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	result = removeElement(nums, 2)
	if result != 5 {
		t.Errorf("Get %d, Expect 5", result)
	}

	nums = []int{0, 4, 4, 0, 4, 4, 4, 0, 2}
	result = removeElement(nums, 4)
	if result != 4 {
		t.Errorf("Get %d, Expect 4", result)
	}

	nums = []int{3, 3}
	result = removeElement(nums, 3)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}
}
