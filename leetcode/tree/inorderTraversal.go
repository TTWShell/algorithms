/* https://leetcode.com/problems/binary-tree-inorder-traversal/description/
Given a binary tree, return the inorder traversal of its nodes' values.

For example:
Given binary tree [1,null,2,3],
   1
    \
     2
    /
   3
return [1,3,2].

Note: Recursive solution is trivial, could you do it iteratively?
*/

package ltree

import "github.com/TTWShell/algorithms/stack" // need copy stack.go when run in leetcode online

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	s := stack.Constructor()

	for root != nil || !s.IsEmpty() {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		root = s.Pop().(*TreeNode)
		res = append(res, root.Val)
		root = root.Right
	}

	return res
}

/* 0ms
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	return append(append(left, root.Val), right...)
}
*/

/* 0ms
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	var inorder func(root *TreeNode, res *[]int)
	inorder = func(root *TreeNode, res *[]int) {
		if root != nil {
			inorder(root.Left, res)
			*res = append(*res, root.Val)
			inorder(root.Right, res)
		}
	}

	inorder(root, &res)
	return res
}
*/
