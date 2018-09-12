package list

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	var l *ListNode
	var result, expect []int

	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	l = l.Reverse()
	result = l.Print()
	expect = []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
