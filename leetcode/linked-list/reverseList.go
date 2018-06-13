/* https://leetcode.com/problems/reverse-linked-list/#/description
Reverse a singly linked list.

Hint:
A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

package lll

func reverseList(head *ListNode) *ListNode {
	var top *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = top
		top = head
		head = tmp
	}
	return top
}
