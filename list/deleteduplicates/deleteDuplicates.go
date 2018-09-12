package deleteduplicates

import "github.com/alandtsang/leetcode-go/list"

type ListNode = list.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	val := head.Val
	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val != val {
			val = p.Next.Val
			p = p.Next
		} else {
			n := p.Next.Next
			p.Next = n
		}
	}
	return head
}
