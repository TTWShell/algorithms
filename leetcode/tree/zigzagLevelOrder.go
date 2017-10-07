/* https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
    Given binary tree [3,9,20,null,null,15,7],
        3
       / \
      9  20
        /  \
       15   7
    return its zigzag level order traversal as:
    [
      [3],
      [20,9],
      [15,7]
    ]
*/

package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	stack, reverse := []*TreeNode{root}, false
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

		if reverse {
			for i := 0; i < len(tmpRes)/2; i++ {
				tmpRes[i], tmpRes[len(tmpRes)-1-i] = tmpRes[len(tmpRes)-1-i], tmpRes[i]
			}
		}

		res, stack, reverse = append(res, tmpRes), tmpStack, !reverse
	}
	return res
}
