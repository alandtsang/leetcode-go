package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	var a, b, ret string

	a = "11"
	b = "1"
	ret = addBinary(a, b)
	if ret != "100" {
		t.Errorf("Get %s, Expect 100", ret)
	}

	a = "1010"
	b = "1011"
	ret = addBinary(a, b)
	if ret != "10101" {
		t.Errorf("Get %s, Expect 10101", ret)
	}
}
