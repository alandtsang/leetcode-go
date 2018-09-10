package reverse

import "testing"

func TestReverse(t *testing.T) {
	var result int

	result = reverse(123)
	if result != 321 {
		t.Errorf("Get %d, Expect 321", result)
	}

	result = reverse(-123)
	if result != -321 {
		t.Errorf("Get %d, Expect -321", result)
	}

	result = reverse(120)
	if result != 21 {
		t.Errorf("Get %d, Expect 21", result)
	}

	result = reverse(0)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}

	result = reverse(2147483647)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}

	result = reverse(-2147483647)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}
}

func TestReverse2(t *testing.T) {
	var result int

	result = reverse2(123)
	if result != 321 {
		t.Errorf("Get %d, Expect 321", result)
	}

	result = reverse2(-123)
	if result != -321 {
		t.Errorf("Get %d, Expect -321", result)
	}

	result = reverse2(120)
	if result != 21 {
		t.Errorf("Get %d, Expect 21", result)
	}

	result = reverse2(0)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}

	result = reverse2(2147483647)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}

	result = reverse2(-2147483647)
	if result != 0 {
		t.Errorf("Get %d, Expect 0", result)
	}
}
