package myhashset

import "testing"

func TestMyHashSet(t *testing.T) {
	var ok bool
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	if ok = hashSet.Contains(1); ok != true { // 返回 true
		t.Errorf("hashSet.Contains(1), get %v, expect true", ok)
	}
	if ok = hashSet.Contains(3); ok != false { // 返回 false（未找到）
		t.Errorf("hashSet.Contains(3), get %v, expect false", ok)
	}
	hashSet.Add(2)
	if ok = hashSet.Contains(2); ok != true { // 返回 true
		t.Errorf("hashSet.Contains(2), get %v, expect true", ok)
	}
	hashSet.Remove(2)
	if ok = hashSet.Contains(2); ok != false { // 返回  false (已经被删除)
		t.Errorf("hashSet.Contains(2), get %v, expect false", ok)
	}
}
