package ismonotonic

import "testing"

func TestIsMonotonic(t *testing.T) {
	var A []int
	var result bool

	A = []int{1, 2, 2, 3}
	result = isMonotonic(A)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	A = []int{6, 5, 4, 4}
	result = isMonotonic(A)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	A = []int{1, 3, 2}
	result = isMonotonic(A)
	if result != false {
		t.Errorf("Got %v, Expect false", result)
	}

	A = []int{1, 2, 4, 5}
	result = isMonotonic(A)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}

	A = []int{1, 1, 1}
	result = isMonotonic(A)
	if result != true {
		t.Errorf("Got %v, Expect true", result)
	}
}
