package insertionsortlist

import (
	"reflect"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	var l *ListNode
	var result, expect []int

	l = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	l = insertionSortList(l)
	result = l.Print()
	expect = []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}
}
