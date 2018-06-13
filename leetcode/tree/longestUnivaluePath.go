/* https://leetcode.com/problems/longest-univalue-path/description/
Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.

Note: The length of path between two nodes is represented by the number of edges between them.

Example 1:

Input:

              5
             / \
            4   5
           / \   \
          1   1   5
Output:

2
Example 2:

Input:

              1
             / \
            4   5
           / \   \
          4   4   5
Output:

2
Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.
*/

package ltree

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var helper func(root *TreeNode, pVal int) int
	helper = func(root *TreeNode, pVal int) int {
		if root == nil || root.Val != pVal {
			return 0
		}
		return 1 + max(helper(root.Left, pVal), helper(root.Right, pVal))
	}

	return max(helper(root.Left, root.Val)+helper(root.Right, root.Val), max(longestUnivaluePath(root.Left), longestUnivaluePath(root.Right)))
}
