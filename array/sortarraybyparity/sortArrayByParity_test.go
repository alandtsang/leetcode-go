package sortarraybyparity

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	var result, expect []int

	result = sortArrayByParity([]int{3, 1, 2, 4})
	expect = []int{4, 2, 1, 3}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	result = sortArrayByParity([]int{3})
	expect = []int{3}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
