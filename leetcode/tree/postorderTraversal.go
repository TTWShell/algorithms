/* https://leetcode.com/problems/binary-tree-postorder-traversal/description/
Given a binary tree, return the postorder traversal of its nodes' values.

For example:
Given binary tree {1,#,2,3},
   1
    \
     2
    /
   3
return [3,2,1].

Note: Recursive solution is trivial, could you do it iteratively?
*/

package ltree

/*
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    res = append(res, postorderTraversal(root.Left)...)
    res = append(res, postorderTraversal(root.Right)...)
    res = append(res, root.Val)
    return res
}
*/

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			res = append([]int{root.Val}, res...)
			root = root.Right
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = top.Left
	}
	return res
}
