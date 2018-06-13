/* https://leetcode.com/problems/missing-number/#/description
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

For example,
Given nums = [0, 1, 3] return 2.

Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

package lmath

func missingNumber(nums []int) int {
	// 注意题目中说的是find the one that is missing from the array 也就是只会缺一个数字
	t := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		t -= nums[i]
	}
	return t
}
