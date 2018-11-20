// 首先链表非空，且非负，数字逆序存储
// 从头部依次遍历链表，将元素相加，保存进位，
// 如果链表长度不同，记得处理多出的元素。
package addtwonumbers

import (
	"github.com/alandtsang/leetcode-go/list"
)

type ListNode = list.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var l ListNode
	head := &l
	var carry, sum int
	for l1 != nil && l2 != nil {
		//fmt.Printf("l1=%d, l2=%d, carry=%d\n", l1.Val, l2.Val, carry)
		sum = l1.Val + l2.Val + carry
		if sum >= 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}
		head.Next = &ListNode{Val: sum}
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil { // len(l1) > len(l2)
		if carry != 0 {
			sum = l1.Val + carry
			if sum >= 10 {
				carry = sum / 10
				sum = sum % 10
			} else {
				carry = 0
			}
			head.Next = &ListNode{Val: sum}
			head = head.Next
			l1 = l1.Next
		} else {
			head.Next = l1
			break
		}
	}
	for l2 != nil { // len(l2) > len(l1)
		if carry != 0 {
			sum = l2.Val + carry
			if sum >= 10 {
				carry = sum / 10
				sum = sum % 10
			} else {
				carry = 0
			}
			head.Next = &ListNode{Val: sum}
			head = head.Next
			l2 = l2.Next
		} else {
			head.Next = l2
			break
		}
	}
	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}
	return l.Next
}
