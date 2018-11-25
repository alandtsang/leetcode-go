package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers2(t *testing.T) {
	tables := []Tables{
		{&ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2}}},
			&ListNode{Val: 4, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5}}},
			[]int{8, 0, 7},
		},
		{&ListNode{Val: 1, Next: &ListNode{Val: 0}},
			&ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0}}},
			[]int{2, 2, 0},
		},
		{&ListNode{Val: 1},
			&ListNode{Val: 9, Next: &ListNode{Val: 9}},
			[]int{1, 0, 0},
		},
	}

	var l *ListNode
	var result []int
	for _, v := range tables {
		l = addTwoNumbers2(v.l1, v.l2)
		if result = l.Print(); reflect.DeepEqual(result, v.expect) != true {
			t.Errorf("l1=%v, l2=%v, get %v, expect %v", v.l1.Print(), v.l2.Print(), result, v.expect)
		}
	}
}
