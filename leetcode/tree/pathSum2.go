/* https://leetcode.com/problems/path-sum-ii/description/
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
    Given the below binary tree and sum = 22,
                  5
                 / \
                4   8
               /   / \
              11  13  4
             /  \    / \
            7    2  5   1
    return
    [
       [5,4,11,2],
       [5,8,4,5]
    ]
*/
package ltree

func pathSum2(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return [][]int{{root.Val}}
	}

	for _, sub := range pathSum2(root.Left, sum-root.Val) {
		res = append(res, append([]int{root.Val}, sub...))
	}
	for _, sub := range pathSum2(root.Right, sum-root.Val) {
		res = append(res, append([]int{root.Val}, sub...))
	}
	return res
}
