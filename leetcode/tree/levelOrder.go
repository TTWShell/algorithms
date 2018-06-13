/* https://leetcode.com/problems/binary-tree-level-order-traversal/description/
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
    Given binary tree [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
    return its level order traversal as:
    [
      [3],
      [9,20],
      [15,7]
    ]
*/

package ltree

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		tmpRes, tmpStack := []int{}, []*TreeNode{}
		for _, node := range stack {
			if node.Left != nil {
				tmpStack = append(tmpStack, node.Left)
			}
			if node.Right != nil {
				tmpStack = append(tmpStack, node.Right)
			}
			tmpRes = append(tmpRes, node.Val)
		}
		res = append(res, tmpRes)
		stack = tmpStack
	}
	return res
}
