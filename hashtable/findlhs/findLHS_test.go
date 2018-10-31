package findlhs

import "testing"

type Tables struct {
	nums   []int
	expect int
}

func TestFindLHS(t *testing.T) {
	var result int
	tables := []Tables{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
	}

	for _, v := range tables {
		if result = findLHS(v.nums); result != v.expect {
			t.Errorf("%v, get %d, expect %d", v.nums, result, v.expect)
		}
	}
}
