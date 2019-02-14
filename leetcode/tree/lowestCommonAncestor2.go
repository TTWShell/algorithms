/* https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia:
“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]

https://assets.leetcode.com/uploads/2018/12/14/binarytree.png

Example 1:

	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	Output: 3
	Explanation: The LCA of nodes 5 and 1 is 3.

Example 2:

	Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
	Output: 5
	Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

Note:

	All of the nodes' values will be unique.
	p and q are different and both values will exist in the binary tree.
*/

package ltree

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var helper func(root *TreeNode, targetVal int, ancestors *[]*TreeNode) bool
	helper = func(root *TreeNode, targetVal int, ancestors *[]*TreeNode) bool {
		if root == nil {
			return false
		}
		if root.Val == targetVal {
			*ancestors = append(*ancestors, root)
			return true
		}
		if (root.Left != nil && helper(root.Left, targetVal, ancestors)) || (root.Right != nil && helper(root.Right, targetVal, ancestors)) {
			*ancestors = append(*ancestors, root)
			return true
		}
		return false
	}

	ancestorsP, ancestorsQ := []*TreeNode{}, []*TreeNode{}
	helper(root, p.Val, &ancestorsP)
	helper(root, q.Val, &ancestorsQ)

	var idx int
	for idx = 0; idx < len(ancestorsP) && idx < len(ancestorsQ); idx++ {
		if ancestorsP[len(ancestorsP)-1-idx] != ancestorsQ[len(ancestorsQ)-1-idx] {
			break
		}
	}
	return ancestorsP[len(ancestorsP)-idx] // if panic, input err
}
