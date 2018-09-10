package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	var nums []int
	var ret int

	nums = []int{2, 2, 1}
	ret = singleNumber(nums)
	if ret != 1 {
		t.Errorf("Get %d, Expect 1", ret)
	}

	nums = []int{4, 1, 2, 1, 2}
	ret = singleNumber(nums)
	if ret != 4 {
		t.Errorf("Get %d, Expect 4", ret)
	}
}
