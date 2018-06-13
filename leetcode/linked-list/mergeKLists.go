/* https://leetcode.com/problems/merge-k-sorted-lists/#/description
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*/

package lll

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var helper func(lists []*ListNode) *ListNode
	helper = func(lists []*ListNode) *ListNode {
		length := len(lists)
		half := length / 2

		if length == 1 {
			return lists[0]
		} else if length == 2 {
			var (
				c1, c2   = lists[0], lists[1]
				res, cur *ListNode
			)

			if c1 == nil {
				return c2
			}
			if c2 == nil {
				return c1
			}

			if c1.Val < c2.Val {
				res, cur, c1 = c1, c1, c1.Next
			} else {
				res, cur, c2 = c2, c2, c2.Next
			}

			for c1 != nil && c2 != nil {
				if c1.Val < c2.Val {
					cur.Next, c1 = c1, c1.Next
				} else {
					cur.Next, c2 = c2, c2.Next
				}
				cur = cur.Next
			}

			if c1 != nil {
				cur.Next = c1
			}
			if c2 != nil {
				cur.Next = c2
			}
			return res
		}
		return mergeKLists([]*ListNode{mergeKLists(lists[half:]), mergeKLists(lists[:half])})
	}

	return helper(lists)
}

/*
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
*/
