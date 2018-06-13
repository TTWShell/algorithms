/* https://leetcode.com/problems/average-of-levels-in-binary-tree/#/description
Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

Example 1:
    Input:
        3
       / \
      9  20
        /  \
       15   7
    Output: [3, 14.5, 11]
    Explanation:
    The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
Note:
The range of node's value is in the range of 32-bit signed integer.
*/

package ltree

func averageOfLevels(root *TreeNode) []float64 {
	maps := make(map[int][]int)

	var helper func(root *TreeNode, level int)
	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		maps[level] = append(maps[level], root.Val)
		helper(root.Left, level+1)
		helper(root.Right, level+1)
	}
	helper(root, 0)

	average := func(nums []int) float64 {
		sum := 0
		for i := range nums {
			sum += nums[i]
		}
		return float64(sum) / float64(len(nums))
	}

	res := make([]float64, len(maps), len(maps))
	for k, v := range maps {
		res[k] = average(v)
	}
	return res
}
