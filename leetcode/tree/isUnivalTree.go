/* https://leetcode.com/problems/univalued-binary-tree/
A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.

Example 1:

	https://assets.leetcode.com/uploads/2018/12/28/unival_bst_1.png
	Input: [1,1,1,1,1,null,1]
	Output: true

Example 2:

	https://assets.leetcode.com/uploads/2018/12/28/unival_bst_2.png
	Input: [2,2,2,5,2]
	Output: false

Note:

	The number of nodes in the given tree will be in the range [1, 100].
	Each node's value will be an integer in the range [0, 99].
*/

package ltree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left != nil && root.Val != root.Left.Val) || (root.Right != nil && root.Val != root.Right.Val) {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
