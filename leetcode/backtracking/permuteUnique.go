/* https://leetcode.com/problems/permutations-ii/description/
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

For example,
[1,1,2] have the following unique permutations:
    [
      [1,1,2],
      [1,2,1],
      [2,1,1]
    ]
*/

package lbacktracking

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	var helper func(nums []int, start int, res *[][]int)
	helper = func(nums []int, start int, res *[][]int) {
		if len(nums)-1 == start {
			*res = append(*res, append([]int{}, nums...))
		}
		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[start] {
				continue
			}
			nums[i], nums[start] = nums[start], nums[i]
			helper(append([]int{}, nums...), start+1, res)
		}
	}

	helper(nums, 0, &res)
	return res
}
