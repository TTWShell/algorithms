/* https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
*/

package leetcode

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		top := stack[len(stack)-1]
		k--
		if k == 0 {
			return top.Val
		}
		stack = stack[:len(stack)-1]
		root = top.Right
	}
	panic("error occurred")
}
