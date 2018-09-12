package reverselist

import (
	"github.com/alandtsang/leetcode-go/list"
)

type ListNode = list.ListNode

func reverseList(head *ListNode) *ListNode {
	return head.Reverse()
}
