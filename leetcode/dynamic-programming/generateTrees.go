/* https://leetcode.com/problems/unique-binary-search-trees-ii/description/
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1...n.

For example,
Given n = 3, your program should return all 5 unique BST's shown below.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

package ldp

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("<%s %d %s>", t.Left, t.Val, t.Right)
}

func generateTrees(n int) []*TreeNode {
	var helper func(nums []int) []*TreeNode
	helper = func(nums []int) []*TreeNode {
		switch len(nums) {
		case 0:
			return []*TreeNode{}
		case 1:
			return []*TreeNode{&TreeNode{Val: nums[0]}}
		default:
			res := []*TreeNode{}
			for i := 0; i < len(nums); i++ {
				lefts, rights := helper(nums[:i]), helper(nums[i+1:])
				if len(lefts) == 0 {
					for _, right := range rights {
						res = append(res, &TreeNode{Val: nums[i], Right: right})
					}
				} else if len(rights) == 0 {
					for _, left := range lefts {
						res = append(res, &TreeNode{Val: nums[i], Left: left})
					}
				} else {
					for _, left := range lefts {
						for _, right := range rights {
							res = append(res, &TreeNode{Val: nums[i], Left: left, Right: right})

						}
					}
				}
			}
			return res
		}
	}

	nums := make([]int, n, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	return helper(nums)
}
