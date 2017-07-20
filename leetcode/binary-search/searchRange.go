/* https://leetcode.com/problems/search-for-a-range/#/description
Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
*/

package leetcode

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for right >= left {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left, right = mid, mid
			for ; left >= 0 && nums[left] == target; left-- {
			}
			for ; right < len(nums) && nums[right] == target; right++ {
			}
			return []int{left + 1, right - 1}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{-1, -1}
}
