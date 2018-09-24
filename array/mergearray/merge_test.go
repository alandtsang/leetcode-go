package mergearray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var nums1, nums2, expect []int

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	expect = []int{1, 2, 2, 3, 5, 6}
	merge(nums1, 3, nums2, 3)
	if reflect.DeepEqual(nums1, expect) != true {
		t.Errorf("Get %v, Expect %v", nums1, expect)
	}

	nums1 = []int{0}
	nums2 = []int{1}
	expect = []int{1}
	merge(nums1, 0, nums2, 1)
	if reflect.DeepEqual(nums1, expect) != true {
		t.Errorf("Get %v, Expect %v", nums1, expect)
	}

	nums1 = []int{2, 0}
	nums2 = []int{1}
	expect = []int{1, 2}
	merge(nums1, 1, nums2, 1)
	if reflect.DeepEqual(nums1, expect) != true {
		t.Errorf("Get %v, Expect %v", nums1, expect)
	}

	nums1 = []int{2, 0, 0, 0}
	nums2 = []int{1, 2, 3}
	expect = []int{1, 2, 2, 3}
	merge(nums1, 1, nums2, 3)
	if reflect.DeepEqual(nums1, expect) != true {
		t.Errorf("Get %v, Expect %v", nums1, expect)
	}
}
