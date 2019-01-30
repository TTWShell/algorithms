/* https://leetcode.com/problems/linked-list-components/
We are given head, the head node of a linked list containing unique integer values.

We are also given the list G, a subset of the values in the linked list.

Return the number of connected components in G, where two values are connected if they appear consecutively in the linked list.

Example 1:

	Input:
	head: 0->1->2->3
	G = [0, 1, 3]
	Output: 2
	Explanation:
	0 and 1 are connected, so [0, 1] and [3] are the two connected components.

Example 2:

	Input:
	head: 0->1->2->3->4
	G = [0, 3, 1, 4]
	Output: 2
	Explanation:
	0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.

Note:

	If N is the length of the linked list given by head, 1 <= N <= 10000.
	The value of each node in the linked list will be in the range [0, N - 1].
	1 <= G.length <= 10000.
	G is a subset of all values in the linked list.
*/

package lll

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	maps := make(map[int]int, len(G))
	for i := range G {
		maps[G[i]] = 0
	}

	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := maps[cur.Val]; ok {
			count++
			for cur = cur.Next; cur != nil; cur = cur.Next {
				if _, ok := maps[cur.Val]; !ok {
					break
				}
			}
			if cur == nil {
				return count
			}
		}
	}
	return count
}
