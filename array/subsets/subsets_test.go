package subsets

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	var result, expect [][]int

	/*result = subsets([]int{1})
	expect = [][]int{{}, {1}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}*/

	/*result = subsets([]int{1, 2})
	expect = [][]int{{}, {1}, {1, 2}, {2}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}*/

	result = subsets([]int{1, 2, 3})
	expect = [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
