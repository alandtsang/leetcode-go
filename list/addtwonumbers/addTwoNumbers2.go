// 如果链表中数字是顺序存储，如果处理？
package addtwonumbers

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var l ListNode
	head := &l
	p := l1.Reverse()
	q := l2.Reverse()
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
	return l.Next.Reverse()
}
