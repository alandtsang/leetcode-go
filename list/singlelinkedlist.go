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
