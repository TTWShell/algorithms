/* https://leetcode.com/problems/balanced-binary-tree/#/description
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
*/

package ltree

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
