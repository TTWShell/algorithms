/* https://leetcode.com/problems/house-robber/#/description
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

d[i] = max(d[i-2]+nums[i], d[i-1])
第i次要么抢，要么不抢，收益总为：第i－2次＋本次 与 i－1 的最大值。
*/

package ldp

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

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
