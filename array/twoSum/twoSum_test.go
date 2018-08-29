package twoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var nums []int
	var target int
	nums = []int{2, 7, 11, 15}
	target = 9

	result := twoSum(nums, target)
	expect := []int{0, 1}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Got: %v, Expect: %v", result, expect)
	}
}
