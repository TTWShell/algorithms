/* https://leetcode.com/problems/insertion-sort-list/description/
Sort a linked list using insertion sort.
*/

package leetcode

func insertionSortList(head *ListNode) *ListNode {
	var res *ListNode
	if head == nil {
		return res
	}

	res = &ListNode{Val: head.Val}
	curHead, curRes := head.Next, res

	for curHead != nil {
		tmp := &ListNode{Val: curHead.Val}
		switch {
		case curHead.Val < curRes.Val:
			if curHead.Val <= res.Val {
				tmp.Next, res = res, tmp
			} else {
				parent, cur := res, res.Next
				for cur.Val < curHead.Val {
					parent, cur = cur, cur.Next
				}
				tmp.Next, parent.Next = cur, tmp
			}
		default:
			curRes.Next = tmp
			curRes = curRes.Next
		}
		curHead = curHead.Next
	}
	return res
}
