package myhashmap

import "testing"

func TestMyHashMap(t *testing.T) {
	var value int
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	if value = hashMap.Get(1); value != 1 {
		t.Errorf("hashMap.Get(1), get %d, expect %d", value, 1)
	}
	if value = hashMap.Get(3); value != -1 {
		t.Errorf("hashMap.Get(3), get %d, expect %d", value, -1)
	}
	hashMap.Put(2, 1)
	if value = hashMap.Get(2); value != 1 {
		t.Errorf("hashMap.Get(2), get %d, expect %d", value, 1)
	}
	hashMap.Remove(2)
	if value = hashMap.Get(2); value != -1 {
		t.Errorf("hashMap.Get(2), get %d, expect %d", value, -1)
	}
}
