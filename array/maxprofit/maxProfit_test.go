package maxprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	var result int

	result = maxProfit([]int{7, 1, 5, 3, 6, 4})
	if result != 7 {
		t.Errorf("Get %d, Expect 7", result)
	}

	result = maxProfit([]int{1, 2, 3, 4, 5})
	if result != 4 {
		t.Errorf("Get %d, Expect 4", result)
	}

	result = maxProfit([]int{7, 6, 4, 3, 1})
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}
}
