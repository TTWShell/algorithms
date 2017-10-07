/* https://leetcode.com/problems/validate-binary-search-tree/description/
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:
    2
   / \
  1   3
Binary tree [2,1,3], return true.
Example 2:
    1
   / \
  2   3
Binary tree [1,2,3], return false.
*/

package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	var helper func(root *TreeNode, min, max int) bool
	helper = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		return min < root.Val && root.Val < max &&
			helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
	}

	return helper(root, math.MinInt64, math.MaxInt64)
}
