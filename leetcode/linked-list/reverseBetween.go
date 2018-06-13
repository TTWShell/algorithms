/* https://leetcode.com/problems/reverse-linked-list-ii/description/
Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,

return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.
*/

package lll

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	// fake front、middle、back head
	frontHead, midHead, backHead := &ListNode{Val: 0}, &ListNode{Val: 0}, &ListNode{Val: 0}
	idx := 1
	// tail of front、middle、back list
	frontTail, midTail, backTail := frontHead, midHead, backHead

	for cur := head; cur != nil; cur, idx = cur.Next, idx+1 {
		tmp := &ListNode{Val: cur.Val}
		switch {
		case idx < m:
			frontTail.Next, frontTail = tmp, tmp
		case m == idx:
			midTail, midHead.Next = tmp, tmp
		case m < idx && idx <= n:
			tmp.Next, midHead.Next = midHead.Next, tmp
		default:
			backTail.Next, backTail = tmp, tmp
		}
	}
	frontTail.Next = midHead.Next
	midTail.Next = backHead.Next
	return frontHead.Next
}
