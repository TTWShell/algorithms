/* https://leetcode.com/problems/house-robber-ii/description/
Note: This is an extension of House Robber.

After robbing those houses on that street, the thief has found himself a new place for his thievery so that he will not get too much attention. This time, all houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, the security system for these houses remain the same as for those in the previous street.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
*/

package ldp

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	robHelper := func(nums []int) int {
		a, b := 0, 0
		for i := 0; i < len(nums); i++ {
			if i%2 == 0 {
				a = max(a+nums[i], b) // 要不抢窃当前i，要不抢窃之前一家（b）
			} else {
				b = max(a, b+nums[i]) // 同上
			}
		}

		return max(a, b)
	}

	return max(robHelper(nums[1:]), robHelper(nums[0:len(nums)-1]))
}
