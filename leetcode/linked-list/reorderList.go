/* https://leetcode.com/problems/reorder-list/description/
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You must do this in-place without altering the nodes' values.

For example,
Given {1,2,3,4}, reorder it to {1,4,2,3}.
*/

package lll

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	tail, reversed, cur := slow, slow, slow.Next
	for cur != nil {
		cur.Next, reversed, cur = reversed, cur, cur.Next
	}
	tail.Next = nil

	idx, cur, curHead, curReversed := 0, head, head, reversed
	for curHead != nil && curHead != tail {
		if idx%2 == 0 {
			curHead = curHead.Next
			cur.Next = curReversed
			curReversed = curReversed.Next
		} else {
			cur.Next = curHead
		}
		cur = cur.Next
		idx++
	}
}
