/* https://leetcode.com/problems/minimum-absolute-difference-in-bst/#/description
Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

Example:

    Input:

       1
        \
         3
        /
       2

    Output:
    1

    Explanation:
        The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).

Note: There are at least two nodes in this BST.
*/

package leetcode

import "math"

func getMinimumDifference(root *TreeNode) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// Given a binary search tree with non-negative values. So use -1 tag prev
	r, prev := math.MaxInt32, -1
	var inorder func(root *TreeNode) int
	inorder = func(root *TreeNode) int {
		if root == nil {
			return r
		}

		inorder(root.Left)

		if prev != -1 {
			r = min(r, root.Val-prev)
		}
		prev = root.Val

		inorder(root.Right)

		return r
	}

	inorder(root)
	return r
}
