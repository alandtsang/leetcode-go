package maxprofit1

import "testing"

func TestMaxProfit(t *testing.T) {
	var result int

	result = maxProfit([]int{7, 1, 5, 3, 6, 4})
	if result != 5 {
		t.Errorf("Get %d, Expect 5", result)
	}

	result = maxProfit([]int{7, 6, 4, 3, 1})
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}

	result = maxProfit([]int{2, 4, 1})
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	result = maxProfit([]int{2, 1, 2, 1, 0, 1, 2})
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}
}
