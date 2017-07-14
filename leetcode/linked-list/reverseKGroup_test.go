package leetcode

import "testing"

func Test_reverseKGroup(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	if r := reverseKGroup(head, 2); r == nil || r.String() != "2 1 4 3 5 <nil>" {
		t.Fatal(r)
	}

	if r := reverseKGroup(head, 2); r == nil || r.String() != "3 2 1 4 5 <nil>" {
		t.Fatal(r)
	}
}
