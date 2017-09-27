/* https://leetcode.com/problems/jump-game/description/
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

For example:
A = [2,3,1,1,4], return true.

A = [3,2,1,0,4], return false.
*/

package leetcode

func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	index, minIndex := n-1, n-1
	for index >= 0 {
		if newIndex := nums[index] + index; newIndex >= minIndex {
			minIndex = index
		}
		index--
	}
	return minIndex == 0
}

/*
func canJump(nums []int) bool {
	dp := make([]bool, len(nums), len(nums))
	if nums[0] > 0 || len(nums) == 1 {
		dp[0] = true
	}
	for i, num := range nums {
		if dp[i] == true {
			for j := i + 1; j < len(nums) && j <= i+num; j++ {
				dp[j] = true
			}
		}
	}
	return dp[len(nums)-1]
}
*/
