package leetcode

import "testing"

func Test_swapPairs(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	if r := swapPairs(head); r == nil || r.String() != "2 1 4 3 <nil>" {
		t.Fatal(r)
	}
}
