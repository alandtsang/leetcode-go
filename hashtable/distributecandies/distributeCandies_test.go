package distributecandies

import "testing"

type Tables struct {
	candies []int
	expect  int
}

func TestDistributeCandies(t *testing.T) {
	var result int
	tables := []Tables{
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{1, 1, 2, 3}, 2},
	}

	for _, v := range tables {
		if result = distributeCandies(v.candies); result != v.expect {
			t.Errorf("%v, get %d, expect %d", v.candies, result, v.expect)
		}
	}
}
