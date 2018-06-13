/* https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
Follow up for "Search in Rotated Sorted Array":
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Write a function to determine if a given target is in the array.

The array may contain duplicates.
*/

package lbs

func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		switch {
		case nums[mid] == target:
			return true
		case nums[start] < nums[mid]: // left rotate
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		case nums[mid] < nums[end]: // right rotate
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		default:
			if len(nums) > 1 {
				return search2(nums[:mid], target) || search2(nums[mid+1:], target)
			}
			return nums[0] == target
		}
	}
	return false
}
