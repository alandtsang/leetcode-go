package istoeplitzmatrix

import "testing"

func TestIsToeplitzMatrix(t *testing.T) {
	var result bool

	result = isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}})
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}
}
