/* https://leetcode.com/problems/linked-list-cycle/description/
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list,
we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
If pos is -1, then there is no cycle in the linked list.

Example 1:

	Input: head = [3,2,0,-4], pos = 1
	Output: true
	Explanation: There is a cycle in the linked list, where tail connects to the second node.

	https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png

Example 2:

	Input: head = [1,2], pos = 0
	Output: true
	Explanation: There is a cycle in the linked list, where tail connects to the first node.
	https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png

Example 3:

	Input: head = [1], pos = -1
	Output: false
	Explanation: There is no cycle in the linked list.
	https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png

Follow up:

	Can you solve it using O(1) (i.e. constant) memory?
*/

package lll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && slow.Next != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow != nil && slow == fast {
			return true
		}
	}
	return false
}
