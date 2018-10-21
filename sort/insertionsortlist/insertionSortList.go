// prev    cur     next
// h   ->  head  ->  4  ->  2  ->  1  ->  3
//         p
package insertionsortlist

import "github.com/alandtsang/leetcode-go/list"

type ListNode = list.ListNode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var h ListNode
	var p, prev, cur, next *ListNode
	h.Next = head // 哨兵节点，指向 head 链表
	cur = head
	next = cur.Next
	for next != nil {
		prev = &h
		p = h.Next
		for p != next && p.Val < next.Val {
			p = p.Next
			prev = prev.Next
		}
		if p == next { // 上步 for 循环 p.Val < next.Val
			cur = next
		} else { // 上步 for 循环 p.Val > next.Val
			cur.Next = next.Next // 先保存 next 后面的链表，以免断掉
			next.Next = p
			prev.Next = next
		}
		next = cur.Next
	}
	return h.Next
}
