/* https://leetcode.com/problems/linked-list-cycle-ii/
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
To represent a cycle in the given linked list,
we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
 If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

Example 1:

	Input: head = [3,2,0,-4], pos = 1
	Output: tail connects to node index 1
	Explanation: There is a cycle in the linked list, where tail connects to the second node.
	https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png

Example 2:

	Input: head = [1,2], pos = 0
	Output: tail connects to node index 0
	Explanation: There is a cycle in the linked list, where tail connects to the first node.
	https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png

Example 3:

	Input: head = [1], pos = -1
	Output: no cycle
	Explanation: There is no cycle in the linked list.
	https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png


Follow up:
	Can you solve it without using extra space?
*/

package lll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	for loop := head; loop != nil; loop = loop.Next {
		slow, fast := loop, loop
		for fast != nil && fast.Next != nil {
			slow, fast = slow.Next, fast.Next.Next
			if slow == loop || fast == loop {
				return loop
			}
			if fast == nil || slow == fast {
				break
			}
		}
	}
	return nil
}
