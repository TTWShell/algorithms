/* https://leetcode.com/problems/diameter-of-binary-tree/#/description
Given a binary tree, you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

Example:
    Given a binary tree
          1
         / \
        2   3
       / \
      4   5
    Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note: The length of path between two nodes is represented by the number of edges between them.
*/

package ltree

func diameterOfBinaryTree(root *TreeNode) int {
	// 使用匿名函数在这里仅仅是为了方便隔离namespace
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}

	diameter := maxDepth(root.Left) + maxDepth(root.Right)

	return max(diameter, max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)))
}
