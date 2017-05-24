/* https://leetcode.com/problems/single-number/#/description
Given an array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
*/

package leetcode

func singleNumber(nums []int) int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] += 1
	}
	for k, v := range maps {
		if v == 1 {
			return k
		}
	}
	panic("input is error: not every element appears twice except for one")
}
