/* https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.
*/

package lll

func deleteDuplicates2(head *ListNode) *ListNode {
	fakeHead := &ListNode{0, head}
	for cur := fakeHead; cur != nil; cur = cur.Next {
		for cur.Next != nil && cur.Next.Next != nil && cur.Next.Val == cur.Next.Next.Val {
			duplicateVal := cur.Next.Val
			// ignore duplicateVal
			for cur.Next != nil && cur.Next.Val == duplicateVal {
				// find no-repeat Next.Val
				cur.Next = cur.Next.Next
			}
		}
	}
	return fakeHead.Next
}

/*
func deleteDuplicates2(head *ListNode) *ListNode {
	var res, resCur *ListNode
	cur, curCount := head, 1
	for cur != nil {
		if cur.Next == nil || cur.Val != cur.Next.Val {
			if curCount == 1 {
				tmp := &ListNode{Val: cur.Val}
				if res == nil {
					res, resCur = tmp, tmp
				} else {
					resCur.Next, resCur = tmp, tmp
				}
			}
			curCount = 0
		}
		curCount++
		cur = cur.Next
	}
	return res
}
*/
