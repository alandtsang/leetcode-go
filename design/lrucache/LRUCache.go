// LRU 先进先出，适合使用队列，
// put 时放在队尾，一旦超出限制长度，队头出队，
// get 时更新元素到队尾。
// O(1) 时间复杂度完成 get 和 put 只能通过 map 存储，
// 队列如果增加和删除需要 O(1) 时间复杂度，需要使用双向链表。
package lrucache

import (
	"container/list"
)

type Node struct {
	key   int
	value int
}

type LRUCache struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:        list.New(),
		m:        make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.l.MoveToBack(e)
		return e.Value.(Node).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.m[key]; ok {
		e.Value = Node{key, value}
		this.l.MoveToBack(e)
	} else {
		if this.l.Len() >= this.capacity {
			delete(this.m, this.l.Front().Value.(Node).key)
			this.l.Remove(this.l.Front())
		}
		this.l.PushBack(Node{key, value})
		this.m[key] = this.l.Back()
	}
}
