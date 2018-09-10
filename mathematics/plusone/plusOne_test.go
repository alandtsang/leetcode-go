package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var digits []int
	var result []int
	var expect []int

	digits = []int{1, 2, 3}
	result = plusOne(digits)
	expect = []int{1, 2, 4}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	digits = []int{4, 3, 2, 1}
	result = plusOne(digits)
	expect = []int{4, 3, 2, 2}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	digits = []int{9}
	result = plusOne(digits)
	expect = []int{1, 0}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	digits = []int{9, 9}
	result = plusOne(digits)
	expect = []int{1, 0, 0}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
