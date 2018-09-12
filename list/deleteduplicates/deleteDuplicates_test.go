package deleteduplicates

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	var head *ListNode
	var result, expect []int

	head = &ListNode{}
	head = deleteDuplicates(head)
	result = head.Print()
	expect = []int{0}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	head = &ListNode{
		Val: 1,
	}
	head = deleteDuplicates(head)
	result = head.Print()
	expect = []int{1}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	head = deleteDuplicates(head)
	result = head.Print()
	expect = []int{1, 2}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	head = deleteDuplicates(head)
	result = head.Print()
	expect = []int{1, 2, 3}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Get %v, Expect %v", result, expect)
	}

}
