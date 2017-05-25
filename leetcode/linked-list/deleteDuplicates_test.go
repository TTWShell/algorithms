package leetcode

import "testing"

func Test_deleteDuplicates(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}}
	if result := deleteDuplicates(input); result.String() != "1 2 3 <nil>" {
		t.Error(result, "1 2 3 <nil>")
	}
}
