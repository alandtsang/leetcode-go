// 使用 bool slice 实现
package myhashset

type MyHashSet struct {
	s []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	s := make([]bool, 1000000)
	for i := range s {
		s[i] = false
	}
	return MyHashSet{
		s: s,
	}
}

func (this *MyHashSet) Add(key int) {
	this.s[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.s[key] = false
}

/** Returns true if this set did not already contain the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.s[key] == true
}
