package matrixreshape

import (
	"reflect"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	var result, expect [][]int

	result = matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4)
	expect = [][]int{{1, 2, 3, 4}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	result = matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4)
	expect = [][]int{{1, 2}, {3, 4}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	result = matrixReshape([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}, 4, 2)
	expect = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	if reflect.DeepEqual(result, expect) != true {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
