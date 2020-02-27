/* https://leetcode.com/problems/binary-tree-maximum-path-sum/
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes
from some starting node to any node in the tree along the
parent-child connections. The path must contain at least
one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6

Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
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

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, helper(root.Left))
		right := max(0, helper(root.Right))

		maxSum = max(maxSum, root.Val+left+right)
		return max(root.Val+left, root.Val+right)
	}

	helper(root)
	return maxSum
}

// func maxPathSum(root *TreeNode) int {
// 	maxSum := root.Val
// 	cache := map[*TreeNode]int{}

// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}
// 		return b
// 	}

// 	var helper func(root *TreeNode)
// 	helper = func(root *TreeNode) {
// 		if root.Left == nil && root.Right == nil {
// 			cache[root] = root.Val
// 			return
// 		}

// 		if root.Left != nil {
// 			helper(root.Left)
// 		}
// 		if root.Right != nil {
// 			helper(root.Right)
// 		}

// 		maxChild, sumPath := math.MinInt64, root.Val
// 		if left, ok := cache[root.Left]; ok {
// 			maxChild = left
// 			sumPath += left
// 		}
// 		if right, ok := cache[root.Right]; ok {
// 			maxChild = max(maxChild, right)
// 			sumPath += right
// 		}

// 		tmp := max(root.Val, root.Val+maxChild)
// 		cache[root] = tmp
// 		maxSum = max(maxSum, max(tmp, max(maxChild, sumPath)))
// 	}

// 	helper(root)
// 	return maxSum
// }
