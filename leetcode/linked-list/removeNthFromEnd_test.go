package lll

import "testing"

func Test_removeNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	if r := removeNthFromEnd(head, 2); r == nil || r.String() != "1 2 3 5 <nil>" {
		t.Fatal("head =", head, "r =", r)
	}

	head = &ListNode{Val: 1}
	if r := removeNthFromEnd(head, 1); r != nil {
		t.Fatal("head =", head, "r =", r)
	}

	head = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	if r := removeNthFromEnd(head, 2); r == nil || r.String() != "2 <nil>" {
		t.Fatal("head =", head, "r =", r)
	}

	if r := removeNthFromEnd(head, 1); r == nil || r.String() != "1 <nil>" {
		t.Fatal("head =", head, "r =", r)
	}

	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	if r := removeNthFromEnd(head, 1); r == nil || r.String() != "1 2 <nil>" {
		t.Fatal("head =", head, "r =", r)
	}
}
