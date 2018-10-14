package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	var cache LRUCache
	var v int

	cache = Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	v = cache.Get(1)
	if v != 1 {
		t.Errorf("Get %d, Expect 1", v)
	}
	cache.Put(3, 3)
	v = cache.Get(2)
	if v != -1 {
		t.Errorf("Get %d, Expect -1", v)
	}
	cache.Put(4, 4)
	v = cache.Get(1)
	if v != -1 {
		t.Errorf("Get %d, Expect -1", v)
	}
	v = cache.Get(3)
	if v != 3 {
		t.Errorf("Get %d, Expect 3", v)
	}
	v = cache.Get(4)
	if v != 4 {
		t.Errorf("Get %d, Expect 4", v)
	}

	cache = Constructor(1)
	cache.Put(2, 1)
	v = cache.Get(2)
	if v != 1 {
		t.Errorf("Get %d, Expect 1", v)
	}
	cache.Put(3, 2)
	v = cache.Get(2)
	if v != -1 {
		t.Errorf("Get %d, Expect -1", v)
	}
	v = cache.Get(3)
	if v != 2 {
		t.Errorf("Get %d, Expect 2", v)
	}
}
