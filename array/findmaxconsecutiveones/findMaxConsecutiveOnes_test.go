package findmaxconsecutiveones

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	var result int

	result = findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	if result != 3 {
		t.Errorf("Get %d, Expect 3", result)
	}
}
