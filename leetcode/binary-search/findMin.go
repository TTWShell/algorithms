/* https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Find the minimum element.

You may assume no duplicate exists in the array.
*/

package leetcode

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	if nums[end] >= nums[start] {
		return nums[start]
	}

	for end-start > 1 {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid
		} else {
			end = mid
		}
	}

	return nums[end]
}
