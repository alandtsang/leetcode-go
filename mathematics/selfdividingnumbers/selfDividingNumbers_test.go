package selfdividingnumbers

import (
	"reflect"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	var result, expect []int

	result = selfDividingNumbers(1, 22)
	expect = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
