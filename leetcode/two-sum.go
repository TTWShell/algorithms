/* https://leetcode.com/problems/two-sum/#/description
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package leetcode

func twoSum(nums []int, target int) []int {
	length := len(nums)
	m := make(map[int]int, length)
	for i := 0; i < length; i++ {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	panic("no two sum solution")
}
