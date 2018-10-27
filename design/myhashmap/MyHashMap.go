// map 的 Get() 是 O(1) 时间复杂度，
// 容易想到通过 slice 及其下标来实现
package myhashmap

type MyHashMap struct {
	m []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	m := make([]int, 10000000)
	for i := range m {
		m[i] = -1
	}
	return MyHashMap{
		m: m,
	}
}

/** value will always be positive. */
func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.m[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.m[key] = -1
}
