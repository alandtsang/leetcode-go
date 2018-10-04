package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	var result int

	result = majorityElement([]int{3, 2, 3})
	if result != 3 {
		t.Errorf("Get %d, Expect 3", result)
	}

	result = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}
}
