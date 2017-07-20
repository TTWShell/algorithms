/* https://leetcode.com/problems/search-in-rotated-sorted-array/#/description
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.
*/

package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left != right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return -1
}
