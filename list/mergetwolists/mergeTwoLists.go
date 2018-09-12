package mergetwolists

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result ListNode
	var head *ListNode
	head = &result

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}
	return result.Next
}
