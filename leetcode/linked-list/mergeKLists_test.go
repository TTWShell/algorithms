package leetcode

import "testing"

func Test_mergeKLists(t *testing.T) {
	lists := []*ListNode{
		&ListNode{Val: 1, Next: &ListNode{Val: 100}},
		&ListNode{Val: 5, Next: &ListNode{Val: 8}},
		&ListNode{Val: 10, Next: &ListNode{Val: 20}},
		&ListNode{Val: 2, Next: &ListNode{Val: 90}},
	}
	if r := mergeKLists(lists); r == nil || r.String() != "1 2 5 8 10 20 90 100 <nil>" {
		t.Fatal(r)
	}
}
