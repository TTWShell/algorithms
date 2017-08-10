/* https://leetcode.com/problems/first-missing-positive/description/
Given an unsorted integer array, find the first missing positive integer.

For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.

Your algorithm should run in O(n) time and uses constant space.
*/

package leetcode

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		index := nums[i] - 1
		if nums[i] < 1 || nums[i] > len(nums) || index < 0 || index > len(nums) || nums[i] == i+1 {
			// cannot swap or noused
			continue
		}
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
