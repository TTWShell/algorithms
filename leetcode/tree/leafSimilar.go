/* https://leetcode.com/problems/leaf-similar-trees/
Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.

https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.
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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var helper func(root *TreeNode, leafs *[]int)
	helper = func(root *TreeNode, leafs *[]int) {
		// helper: find all leaf nodes
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*leafs = append(*leafs, root.Val)
			return
		}
		if root.Left != nil {
			helper(root.Left, leafs)
		}
		if root.Right != nil {
			helper(root.Right, leafs)
		}
	}

	leaf1, leaf2 := []int{}, []int{}
	helper(root1, &leaf1)
	helper(root2, &leaf2)

	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}
