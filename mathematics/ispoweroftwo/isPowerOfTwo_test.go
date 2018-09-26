package ispoweroftwo

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	var result bool

	result = isPowerOfTwo(0)
	if result != false {
		t.Errorf("Get %v, Expect false", result)
	}

	result = isPowerOfTwo(1)
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}

	result = isPowerOfTwo(16)
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}

	result = isPowerOfTwo(218)
	if result != false {
		t.Errorf("Get %v, Expect false", result)
	}
}
