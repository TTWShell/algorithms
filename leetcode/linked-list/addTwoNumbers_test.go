package leetcode

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	if result := addTwoNumbers(l1, l2); result.String() != "7 0 0 1 <nil>" {
		t.Log(result)
		t.Fail()
	}
}
