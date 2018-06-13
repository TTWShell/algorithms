/* https://leetcode.com/problems/single-number-iii/description/
Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

For example:

Given nums = [1, 2, 1, 3, 2, 5], return [3, 5].

Note:
    The order of the result is not important. So in the above example, [5, 3] is also correct.
    Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?
*/

package lbm

func singleNumber3(nums []int) []int {
	diff := 0
	for _, num := range nums {
		diff ^= num
	}
	// Get its last set bit
	diff &= -diff

	res := make([]int, 2, 2)
	for _, num := range nums {
		if num&diff == 0 { // the bit is not select
			res[0] ^= num
		} else { // the bit is set
			res[1] ^= num
		}
	}
	return res
}
