/* https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
Suppose an array sorted in ascending order is rotated at some pivot
unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

The array may contain duplicates.

Example 1:

	Input: [1,3,5]
	Output: 1

Example 2:

	Input: [2,2,2,0,1]
	Output: 0
Note:

This is a follow up problem to Find Minimum in Rotated Sorted Array.
Would allow duplicates affect the run-time complexity? How and why?
*/

package lbs

func findMin2(nums []int) int {
	start, end := 0, len(nums)-1
	if nums[end] > nums[start] {
		return nums[start]
	}

	for end-start > 1 {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid
		} else if nums[mid] == nums[end] {
			cur := mid
			for cur+1 <= end && nums[cur] == nums[cur+1] {
				cur++
			}
			if cur >= end {
				end = mid
			} else {
				start = mid
			}
		} else {
			end = mid
		}
	}
	return nums[end]
}
