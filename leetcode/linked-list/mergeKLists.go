/* https://leetcode.com/problems/merge-k-sorted-lists/#/description
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*/

package leetcode

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		res, tmp      *ListNode
		minIndex, min int
	)

	for cur := res; ; lists[minIndex] = lists[minIndex].Next {
		minIndex, min = -1, math.MaxInt32
		for i, list := range lists {
			if list != nil && list.Val < min {
				min = list.Val
				minIndex = i
			}
		}

		if minIndex == -1 {
			break
		}

		tmp = &ListNode{Val: min}
		if res == nil {
			res = tmp
		} else {
			cur.Next = tmp
		}
		cur = tmp
	}
	return res
}
