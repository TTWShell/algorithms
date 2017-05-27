/* https://leetcode.com/problems/contains-duplicate-ii/#/description
Given an array of integers and an integer k,
find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j]
and the absolute difference between i and j is at most k.

给出一个整数数组，判断该数组内是否有两个元素值是相同的，且他们的索引值相差不大于k，是则返回true，否则返回false
*/

package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	maps := make(map[int]int)
	for i := range nums {
		if _, ok := maps[nums[i]]; !ok {
			maps[nums[i]] = i
		} else {
			if i-maps[nums[i]] <= k {
				return true
			}
			maps[nums[i]] = i
		}
	}
	return false
}
