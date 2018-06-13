/* https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
*/

package ltree

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%d %s", l.Val, l.Next)
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	slow, fast := head, head
	left, pLeft := slow, slow
	for fast != nil && fast.Next != nil {
		pLeft, slow, fast = slow, slow.Next, fast.Next.Next
	}
	right := slow.Next
	pLeft.Next = nil
	return &TreeNode{
		Left:  sortedListToBST(left),
		Val:   slow.Val,
		Right: sortedListToBST(right),
	}
}
