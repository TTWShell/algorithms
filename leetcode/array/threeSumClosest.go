/* https://leetcode.com/problems/3sum-closest/#/description
Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

    For example, given array S = {-1 2 1 -4}, and target = 1.

    The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/

package larray

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	sum := nums[0] + nums[1] + nums[2]
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip same result
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			temp := nums[j] + nums[k] + nums[i]
			if temp == target {
				return target
			} else if temp > target {
				k--
			} else {
				j++
			}
			if abs(temp-target) < abs(sum-target) {
				sum = temp
			}
		}
	}

	return sum
}
