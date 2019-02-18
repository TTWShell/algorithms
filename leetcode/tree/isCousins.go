/* https://leetcode.com/problems/cousins-in-binary-tree/description/
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.

Example 1:

	https://assets.leetcode.com/uploads/2019/02/12/q1248-01.png
	Input: root = [1,2,3,4], x = 4, y = 3
	Output: false

Example 2:

	https://assets.leetcode.com/uploads/2019/02/12/q1248-02.png
	Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
	Output: true

Example 3:

	https://assets.leetcode.com/uploads/2019/02/13/q1248-03.png
	Input: root = [1,2,3,null,4], x = 2, y = 3
	Output: false

Note:

	The number of nodes in the tree will be between 2 and 100.
	Each node has a unique integer value from 1 to 100.
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
func isCousins(root *TreeNode, x int, y int) bool {
	var helper func(root, parent *TreeNode, target, depth int) (p *TreeNode, d int, ok bool)
	helper = func(root, parent *TreeNode, target, depth int) (p *TreeNode, d int, ok bool) {
		if root.Val == target {
			return parent, depth, true
		}
		if root.Left != nil {
			p, d, ok := helper(root.Left, root, target, depth+1)
			if ok {
				return p, d, ok
			}
		}
		if root.Right != nil {
			p, d, ok := helper(root.Right, root, target, depth+1)
			if ok {
				return p, d, ok
			}
		}
		return nil, 0, false
	}

	if root == nil || x == root.Val || y == root.Val {
		return false
	}
	px, dx, _ := helper(root, nil, x, 0)
	py, dy, _ := helper(root, nil, y, 0)
	if dx != dy || px == py {
		return false
	}
	return true
}
