package leetcode

import "testing"

func Test_reverseList(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 6}}}}}
	if r := reverseList(head); r.String() != "6 3 6 2 1 <nil>" {
		t.Error(r)
	}
}
