/* https://leetcode.com/problems/merge-two-sorted-lists/#/description
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
*/
package lll

// ListNode definded in addTwoNumbers
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	c1, c2 := l1, l2
	var c, r, temp *ListNode

	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			temp = &ListNode{Val: c1.Val}
			c1 = c1.Next
		} else {
			temp = &ListNode{Val: c2.Val}
			c2 = c2.Next
		}
		if r == nil {
			r = temp
		} else {
			c.Next = temp
		}
		c = temp
	}

	if r == nil {
		if c1 != nil {
			return c1
		} else {
			return c2
		}
	}

	if c1 != nil {
		c.Next = c1
	}
	if c2 != nil {
		c.Next = c2
	}
	return r
}
