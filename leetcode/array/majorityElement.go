/* https://leetcode.com/problems/majority-element/#/description
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Credits:
Special thanks to @ts for adding this problem and creating all test cases.
*/

package leetcode

func majorityElement(nums []int) int {
	maps := make(map[int]int)
	mid := len(nums) / 2
	for i := 0; i < len(nums); i++ {
		if v, ok := maps[nums[i]]; ok {
			v += 1
			maps[nums[i]] = v
		} else {
			maps[nums[i]] = 1
		}
		if v, ok := maps[nums[i]]; ok && v > mid {
			return nums[i]
		}
	}
	panic("nums cannot be empty")
}
