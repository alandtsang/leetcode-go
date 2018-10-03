package arraypairsum

import "testing"

func TestArrayPairSum(t *testing.T) {
	var result int

	result = arrayPairSum([]int{1, 4, 3, 2})
	if result != 4 {
		t.Errorf("Get %d, Expect 4", result)
	}
}
