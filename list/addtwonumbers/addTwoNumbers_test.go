package addtwonumbers

import (
	"reflect"
	"testing"
)

type Tables struct {
	l1     *ListNode
	l2     *ListNode
	expect []int
}

func TestAddTwoNumbers(t *testing.T) {
	tables := []Tables{
		{&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			[]int{7, 0, 8},
		},
		{&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
			[]int{7, 0, 8, 3},
		},
		{&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 8}}},
			[]int{7, 0, 2, 1},
		},
		{&ListNode{Val: 1},
			&ListNode{Val: 9, Next: &ListNode{Val: 9}},
			[]int{0, 0, 1},
		},
		{&ListNode{Val: 9, Next: &ListNode{Val: 8}},
			&ListNode{Val: 1},
			[]int{0, 9},
		},
	}

	var l *ListNode
	var result []int
	for _, v := range tables {
		l = addTwoNumbers(v.l1, v.l2)
		if result = l.Print(); reflect.DeepEqual(result, v.expect) != true {
			t.Errorf("l1=%v, l2=%v, get %v, expect %v", v.l1.Print(), v.l2.Print(), result, v.expect)
		}

	}
}
