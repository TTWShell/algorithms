/* https://leetcode.com/problems/sort-list/description/
Sort a linked list in O(n log n) time using constant space complexity.
*/

package lll

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// cut List
	slow, fast := head, head
	var tail *ListNode
	for fast != nil && fast.Next != nil {
		tail, slow, fast = slow, slow.Next, fast.Next.Next
	}
	tail.Next = nil

	return sortListHelper(sortList(head), sortList(slow))
}

func sortListHelper(left, right *ListNode) (head *ListNode) {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	var p, cur *ListNode
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur, left = left, left.Next
		} else {
			cur, right = right, right.Next
		}

		if head == nil {
			head = cur
		} else {
			p.Next = cur
		}
		p = cur
	}

	if left != nil {
		p.Next = left
	} else {
		p.Next = right
	}
	return head
}
