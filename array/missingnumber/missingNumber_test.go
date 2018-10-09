package missingnumber

import "testing"

func TestMissingNumber(t *testing.T) {
	var result int

	result = missingNumber([]int{3, 0, 1})
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	result = missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})
	if result != 8 {
		t.Errorf("Get %d, Expect 8", result)
	}
}
