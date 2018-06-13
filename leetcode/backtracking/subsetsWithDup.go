/* https://leetcode.com/problems/subsets-ii/description/
Given a collection of integers that might contain duplicates, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,2], a solution is:

[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

package lbacktracking

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var helper func(nums []int, k int) [][]int
	helper = func(nums []int, k int) [][]int {
		var res [][]int
		switch k {
		case 1:
			res = append(res, []int{nums[0]})
			for i := 1; i < len(nums); i++ {
				if num := nums[i]; res[len(res)-1][0] != num {
					res = append(res, []int{num})
				}
			}
		default:
			for i := 0; i < len(nums)-k+1; i++ {
				if i == 0 || nums[i] != nums[i-1] {
					for _, sub := range helper(nums[i+1:], k-1) {
						res = append(res, append([]int{nums[i]}, sub...))
					}
				}
			}
		}
		return res
	}

	sort.Ints(nums)
	res := [][]int{{}}
	for k := 1; k <= len(nums); k++ {
		res = append(res, helper(nums, k)...)
	}
	return res
}
