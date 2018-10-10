package finddisappearednumbers

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	var result, expect []int

	result = findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	expect = []int{5, 6}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
