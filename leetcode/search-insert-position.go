/* https://leetcode.com/problems/search-insert-position/#/description
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Here are few examples.
[1,3,5,6], 5 → 2
[1,3,5,6], 2 → 1
[1,3,5,6], 7 → 4
[1,3,5,6], 0 → 0
*/

package leetcode

func searchInsert(nums []int, target int) int {
	start, end, mid := 0, len(nums)-1, 0
	if target < nums[start] {
		return 0
	}
	if target > nums[end] {
		return end + 1
	}
	for start <= end {
		mid = start + (end-start)/2
		switch {
		case target == nums[mid]:
			return mid
		case target > nums[mid]:
			start = mid + 1
		case target < nums[mid]:
			end = mid - 1
		}
	}
	return start
}
