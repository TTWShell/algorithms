package leetcode

import "testing"

func Test_removeElements(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 6}}}}}
	if r := removeElements(head, 6); r.String() != "1 2 3 <nil>" {
		t.Fatal(r)
	}
}
