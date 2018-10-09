package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	var result bool

	result = containsDuplicate([]int{1, 2, 3, 1})
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}

	result = containsDuplicate([]int{1, 2, 3, 4})
	if result != false {
		t.Errorf("Get %v, Expect false", result)
	}

	result = containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}
}
