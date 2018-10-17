package adddigits

import "testing"

func TestAddDigits(t *testing.T) {
	var result int

	if result = addDigits(38); result != 2 {
		t.Errorf("Get %d, Expect 2", result)
	}

	if result = addDigits(12); result != 3 {
		t.Errorf("Get %d, Expect 3", result)
	}

	if result = addDigits(1234); result != 1 {
		t.Errorf("Get %d, Expect 1", result)
	}

	if result = addDigits(1234567); result != 1 {
		t.Errorf("Get %d, Expect 1", result)
	}
}
