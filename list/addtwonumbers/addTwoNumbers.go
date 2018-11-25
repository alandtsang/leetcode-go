// 首先链表非空，且非负，数字逆序存储
// 从头部依次遍历链表，将元素相加，保存进位，
// 如果遍历完仍有进位，记得处理。
package addtwonumbers

import (
	"github.com/alandtsang/leetcode-go/list"
)

type ListNode = list.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l ListNode
	head := &l
	p := l1
	q := l2
	var carry, sum, x, y int
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum = x + y + carry
		carry = sum / 10
		head.Next = &ListNode{Val: sum % 10}
		head = head.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		head.Next = &ListNode{Val: carry}
	}
	return l.Next
}
