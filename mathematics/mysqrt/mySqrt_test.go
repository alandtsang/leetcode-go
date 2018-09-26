package mysqrt

import "testing"

func TestMySqrt(t *testing.T) {
	var result int

	result = mySqrt(1)
	if result != 1 {
		t.Errorf("Get %d, Expect 1", result)
	}

	result = mySqrt(4)
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	result = mySqrt(6)
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	result = mySqrt(8)
	if result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}
}
