/* https://leetcode.com/problems/remove-linked-list-elements/#/description
Remove all elements from a linked list of integers that have value val.

Example
Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
Return: 1 --> 2 --> 3 --> 4 --> 5
*/

package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	var r, cur, temp *ListNode
	for ; head != nil; head = head.Next {
		if head.Val != val {
			temp = &ListNode{Val: head.Val}
			if r == nil {
				r = temp
			} else {
				cur.Next = temp
			}
			cur = temp
		}
	}
	return r
}
