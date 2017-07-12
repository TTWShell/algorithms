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

	var res, cur *ListNode
	for i := 0; i < total-n; i++ {
		temp := &ListNode{Val: head.Val}
		if res == nil {
			res = temp
		} else {
			cur.Next = temp

		}
		cur, head = temp, head.Next
	}

	if cur != nil {
		cur.Next = head.Next
	}
	return res
}
