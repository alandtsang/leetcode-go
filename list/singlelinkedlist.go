package list

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() []int {
	var nums []int

	for l != nil {
		nums = append(nums, l.Val)
		l = l.Next
	}
	//fmt.Println(nums)
	return nums
}

func (l *ListNode) Reverse() *ListNode {
	if l == nil {
		return l
	}

	var tail, p, next *ListNode
	p, next = l, l
	for p != nil && next != nil {
		next = p.Next
		p.Next = tail
		tail = p
		p = next
	}
	p = tail // tail 指向 l 最后一个元素
	return p
}
