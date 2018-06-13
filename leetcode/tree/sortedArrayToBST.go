/* https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/#/description
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
*/

package ltree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return toBST(nums, 0, len(nums)-1)
}

func toBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = toBST(nums, start, mid-1)
	root.Right = toBST(nums, mid+1, end)
	return root
}
