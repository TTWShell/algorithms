package leetcode

import "testing"

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	if result := mergeTwoLists(l1, l2); result.String() != "1 2 4 4 5 5 6 <nil>" {
		t.Error(result)
	}

	l1 = nil
	l2 = &ListNode{Val: 1}
	if result := mergeTwoLists(l1, l2); result.String() != "1 <nil>" {
		t.Error(result)
	}
}
