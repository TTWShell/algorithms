/* https://leetcode.com/problems/sum-root-to-leaf-numbers/description/
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

For example,

    1
   / \
  2   3
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.

Return the sum = 12 + 13 = 25.
*/

package ltree

import "math"

func sumNumbers(root *TreeNode) int {
	var helper func(root *TreeNode) (res [][]int)
	helper = func(root *TreeNode) (res [][]int) {
		if root != nil {
			if root.Left == nil && root.Right == nil {
				res = [][]int{{root.Val}}
			} else {
				for _, sub := range helper(root.Left) {
					res = append(res, append(sub, root.Val))
				}
				for _, sub := range helper(root.Right) {
					res = append(res, append(sub, root.Val))
				}
			}
		}
		return res
	}

	sum := 0
	for _, path := range helper(root) {
		for i, digit := range path {
			sum += int(math.Pow10(i)) * digit
		}
	}
	return sum
}
