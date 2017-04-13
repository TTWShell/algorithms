/* https://leetcode.com/problems/add-two-numbers/#/description
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

给你两个表示两个非负数字的链表。数字以相反的顺序存储，其节点包含单个数字。将这两个数字相加并将其作为一个链表返回。

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/
package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%d %s", l.Val, l.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result, current *ListNode
		sum             = 0
		carry           = 0
	)

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10

		temp := &ListNode{Val: sum % 10}
		if result == nil {
			result = temp
		} else {
			current.Next = temp
		}
		current = temp
	}
	return result
}
