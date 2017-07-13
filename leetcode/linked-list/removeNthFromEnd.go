/* https://leetcode.com/problems/remove-nth-node-from-end-of-list/#/description
Given a linked list, remove the n^th node from the end of list and return its head.

For example,

   Given linked list: 1->2->3->4->5, and n = 2.

   After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
    Given n will always be valid.
    Try to do this in one pass.
*/

package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{}
	slow, fast := res, head
	slow.Next = head

	// Move fast in front so that the gap between slow and fast becomes n
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}
	// Move fast to the end, maintaining the gap
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// Skip the desired node
	slow.Next = slow.Next.Next
	return res.Next
}

/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i, total := 0, 0
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		i++
		slow, fast = slow.Next, fast.Next.Next
	}
	if fast == nil {
		total = i * 2
	} else {
		total = i*2 + 1
	}

	if total-n == 0 {
		return head.Next
	}

	var res, cur, ptr *ListNode

	if total-n > i {
		res, cur, ptr = head, slow, slow.Next
		i++
	} else {
		i = 0
		res, ptr = head, head
	}

	for ; i < total-n; i++ {
		cur, ptr = ptr, ptr.Next
	}

	if cur != nil {
		cur.Next = ptr.Next
	}
	return res
}
*/
