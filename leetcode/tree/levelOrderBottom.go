/* https://leetcode.com/problems/binary-tree-level-order-traversal-ii/#/description
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	tempR := [][]int{}
	if root == nil {
		return tempR
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		s := []*TreeNode{}
		r := []int{}
		for i := 0; i < len(stack); i++ {
			r = append(r, stack[i].Val)
			if stack[i].Left != nil {
				s = append(s, stack[i].Left)
			}
			if stack[i].Right != nil {
				s = append(s, stack[i].Right)
			}
		}
		tempR = append(tempR, r)
		stack = s
	}

	n := len(tempR)
	for i := 0; i < n/2; i++ {
		temp := tempR[n-1-i]
		tempR[n-1-i], tempR[i] = tempR[i], temp
	}
	return tempR
}
