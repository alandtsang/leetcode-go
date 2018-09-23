package ispalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	var head *ListNode
	var result bool

	head = &ListNode{
		Val: 1,
	}
	result = isPalindrome(head)
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	result = isPalindrome(head)
	if result != false {
		t.Errorf("Get %v, Expect false", result)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
			},
		},
	}
	result = isPalindrome(head)
	if result != false {
		t.Errorf("Get %v, Expect false", result)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	result = isPalindrome(head)
	if result != true {
		t.Errorf("Get %v, Expect true", result)
	}

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	result = isPalindrome(head)
	if result != false {
		t.Errorf("Get %v, Expect false", result)
	}
}
