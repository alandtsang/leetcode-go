package ispalindrome

import (
	"github.com/alandtsang/leetcode-go/list"
)

type ListNode = list.ListNode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var fast, slow *ListNode // 快慢指针
	fast = head
	slow = head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow
	slow = head
	fast = fast.Reverse() // 右半部链表翻转

	for slow != nil && fast != nil { // 对比左右两个链表
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
