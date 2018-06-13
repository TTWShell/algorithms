package lll

import "testing"

func Test_isPalindrome(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}}

	input := []*ListNode{l1, l2}
	result := []bool{false, true}

	for i := range input {
		if r := isPalindrome(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
