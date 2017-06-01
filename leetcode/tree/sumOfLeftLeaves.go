/* https://leetcode.com/problems/sum-of-left-leaves/#/description
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
*/

package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil {
		if root.Left.Left != nil || root.Left.Right != nil {
			sum += sumOfLeftLeaves(root.Left)
		} else {
			sum += root.Left.Val
		}
	}
	return sum + sumOfLeftLeaves(root.Right)
}
