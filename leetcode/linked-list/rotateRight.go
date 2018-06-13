/* https://leetcode.com/problems/rotate-list/description/
Given a list, rotate the list to the right by k places, where k is non-negative.

For example:
Given 1->2->3->4->5->NULL and k = 2,
return 4->5->1->2->3->NULL.
*/

package lll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n, cur := 1, head
	for cur.Next != nil {
		n++
		cur = cur.Next
	}

	cur.Next = head
	m := n - k%n
	for i := 0; i < m; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead
}
