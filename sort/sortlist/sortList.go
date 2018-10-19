// 要在 O(nlogn) 时间复杂度完成，可以将链表拆分成两个链表，
// 然后进行归并
package sortlist

import (
	"github.com/alandtsang/leetcode-go/list"
)

type ListNode = list.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var fast, slow, prev *ListNode
	fast = head
	slow = head
	prev = head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return merge(sortList(head), sortList(slow))
}

func merge(l1, l2 *ListNode) *ListNode {
	var dummy ListNode
	var cur *ListNode
	cur = &dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
