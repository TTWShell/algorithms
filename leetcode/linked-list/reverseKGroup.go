/* https://leetcode.com/problems/reverse-nodes-in-k-group/#/description
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
*/

package lll

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	var (
		res, temp *ListNode
		cur, root = res, head
	)

	for root != nil {
		i, nextCur, head := 1, &ListNode{Val: root.Val}, root
		root = root.Next
		for ; i < k && root != nil; i++ {
			if i == 1 {
				temp = &ListNode{Val: root.Val, Next: nextCur}
			} else {
				temp = &ListNode{Val: root.Val, Next: temp}
			}
			root = root.Next
		}

		if i == k {
			if res == nil {
				res = temp
			} else {
				cur.Next = temp
			}
			cur = nextCur
		} else if res == nil {
			return head
		} else if head != nil {
			cur.Next = head
			break
		}
	}
	return res
}
