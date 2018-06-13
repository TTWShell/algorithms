/* https://leetcode.com/problems/single-number-ii/description/
Given an array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
*/

package lbm

func singleNumber2(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = (a ^ nums[i]) & ^b
		b = (b ^ nums[i]) & ^a
	}
	return a
}
