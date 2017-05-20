/* https://leetcode.com/problems/same-tree/#/description
Given two binary trees, write a function to check if they are equal or not.

Two binary trees are considered equal if they are structurally identical and the nodes have the same value.
*/
package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("<%s %d %s>", t.Left, t.Val, t.Right)
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return (p == nil && q == nil) ||
		(p != nil && q != nil &&
			p.Val == q.Val &&
			isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right))
}
