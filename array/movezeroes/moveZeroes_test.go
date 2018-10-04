package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var nums, expect []int

	nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	expect = []int{1, 3, 12, 0, 0}
	if reflect.DeepEqual(nums, expect) != true {
		t.Errorf("Get %v, Expect %v", nums, expect)
	}
}
