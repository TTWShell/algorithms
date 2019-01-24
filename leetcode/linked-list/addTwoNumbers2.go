/* https://leetcode.com/problems/add-two-numbers-ii/
You are given two non-empty linked lists representing two non-negative integers.
The most significant digit comes first and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
	What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

	Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 8 -> 0 -> 7
*/
package lll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	list2array := func(head *ListNode) []int {
		res := []int{}
		for p := head; p != nil; p = p.Next {
			res = append(res, p.Val)
		}
		return res
	}

	a1, a2 := list2array(l1), list2array(l2)
	if len(a1) < len(a2) {
		a1, a2 = a2, a1
	}
	diff := len(a1) - len(a2)
	if diff > 0 {
		a2 = append(make([]int, diff, diff), a2...)
	}

	sum, carry := 0, 0
	for idx := len(a2) - 1; idx >= 0; idx-- {
		sum = carry + a1[idx] + a2[idx]
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		a1[idx] = sum
	}

	var (
		res, p, node *ListNode
		idx          = 0
	)
	if carry == 1 {
		res = &ListNode{Val: 1}
	} else {
		res = &ListNode{Val: a1[0]}
		idx++
	}
	p = res

	for ; idx < len(a1); idx++ {
		node = &ListNode{Val: a1[idx]}
		p.Next = node
		p = node
	}
	return res
}
