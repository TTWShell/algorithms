/* https://leetcode.com/problems/remove-duplicates-from-sorted-list/#/description
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

package lll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var r, c, temp *ListNode
	if head != nil {
		c = &ListNode{Val: head.Val}
		r = c
		head = head.Next
	}

	for head != nil {
		if c.Val != head.Val {
			temp = &ListNode{Val: head.Val}
			c.Next = temp
			c = temp
		}
		head = head.Next
	}
	return r
}
