/* https://leetcode.com/problems/swap-nodes-in-pairs/#/description
Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.
*/

package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var res, cur, temp *ListNode

	for ; head != nil && head.Next != nil; head = head.Next.Next {
		temp = &ListNode{Val: head.Next.Val, Next: &ListNode{Val: head.Val}}
		if res == nil {
			res = temp
		} else {
			cur.Next = temp
		}
		cur = temp.Next
	}

	if head != nil {
		cur.Next = head
	}
	return res
}
