/* https://leetcode.com/problems/4sum/#/description
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.

    For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

    A solution set is:
    [
      [-1,  0, 0, 1],
      [-2, -1, 1, 2],
      [-2,  0, 0, 2]
    ]
*/

package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	KSum(nums, target, 4, []int{}, &res)
	return res
}

func KSum(nums []int, target, K int, result []int, results *[][]int) {
	if len(nums) < K || K < 2 {
		return
	}

	if K == 2 {
		// 2-sum
		left, right := 0, len(nums)-1
		for left < right {
			if temp := nums[left] + nums[right]; temp == target {
				*results = append(*results, append(result, nums[left], nums[right]))
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if temp < target {
				left++
			} else {
				right--
			}
		}
	} else {
		for i := 0; i <= len(nums)-K; i++ {
			if target < nums[i]*K || target > nums[len(nums)-1]*K {
				break
			}
			if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
				KSum(nums[i+1:], target-nums[i], K-1, append(result, nums[i]), results)
			}
		}
	}
}
