package mergetwolists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var l1, l2, l3 *ListNode
	var result, expect []int

	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l3 = mergeTwoLists(l1, l2)
	result = l3.Print()
	expect = []int{1, 1, 2, 3, 4, 4}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
