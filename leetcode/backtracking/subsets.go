/* https://leetcode.com/problems/subsets/description/
Given a set of distinct integers, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:

[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

package lbacktracking

func subsets(nums []int) [][]int {
	var helper func(nums []int, k int) [][]int
	helper = func(nums []int, k int) [][]int {
		var res [][]int
		switch k {
		case 1:
			res = make([][]int, len(nums), len(nums))
			for i, num := range nums {
				res[i] = []int{num}
			}
		default:
			for i := 0; i < len(nums)-k+1; i++ {
				for _, sub := range helper(nums[i+1:], k-1) {
					res = append(res, append([]int{nums[i]}, sub...))
				}
			}
		}
		return res
	}

	res := [][]int{{}}
	for k := 1; k <= len(nums); k++ {
		res = append(res, helper(nums, k)...)
	}
	return res
}
