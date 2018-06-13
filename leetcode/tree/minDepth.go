/* https://leetcode.com/problems/minimum-depth-of-binary-tree/#/description
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
给定一棵二叉树，计算其最小深度。
最小深度是指从根节点出发到达最近的叶子节点所需要经过的节点个数。
*/

package ltree

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 子树是null的时候，不是叶子节点
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
