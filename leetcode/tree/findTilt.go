/* https://leetcode.com/problems/binary-tree-tilt/#/description
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
    Input:
             1
           /   \
          2     3
    Output: 1
    Explanation:
        Tilt of node 2 : 0
        Tilt of node 3 : 0
        Tilt of node 1 : |2-3| = 1
        Tilt of binary tree : 0 + 0 + 1 = 1
Note:
    The sum of node values in any subtree won't exceed the range of 32-bit integer.
    All the tilt values won't exceed the range of 32-bit integer.
*/

package ltree

func findTilt(root *TreeNode) int {
	sum := 0

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var postorder func(root *TreeNode, sum *int) int
	postorder = func(root *TreeNode, sum *int) int {
		if root == nil {
			return 0
		}
		leftSum := postorder(root.Left, sum)
		rightSum := postorder(root.Right, sum)
		*sum += abs(leftSum - rightSum)
		return leftSum + rightSum + root.Val
	}

	postorder(root, &sum)
	return sum
}
