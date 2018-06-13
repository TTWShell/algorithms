/* https://leetcode.com/problems/3sum/#/description
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package larray

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip same result
			continue
		}
		j, k := i+1, len(nums)-1
		target := -nums[i]
		for j < k {
			if temp := nums[j] + nums[k]; temp == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++ // skip same result
				}
				for j < k && nums[k] == nums[k+1] {
					k-- // skip same result
				}
			} else if temp > target {
				k--
			} else {
				j++
			}
		}
	}

	return result
}
