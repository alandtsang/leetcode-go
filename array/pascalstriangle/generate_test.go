package pascalstriangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	var result, expect [][]int

	result = generate(1)
	expect = [][]int{{1}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	result = generate(2)
	expect = [][]int{{1}, {1, 1}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	result = generate(3)
	expect = [][]int{{1}, {1, 1}, {1, 2, 1}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	result = generate(5)
	expect = [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
