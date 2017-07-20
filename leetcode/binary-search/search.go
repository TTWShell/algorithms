/* https://leetcode.com/problems/search-in-rotated-sorted-array/#/description
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.
*/

package leetcode

func search(nums []int, target int) int {
	var (
		lenn        = len(nums)
		left, right = 0, lenn - 1
	)

	if lenn == 0 {
		return -1
	}

	if nums[0] > nums[lenn-1] {
		// is Rotated Sorted Array
		if target <= nums[lenn-1] {
			// right block, find left index
			tmp := right
			for tmp >= left {
				mid := left + (tmp-left)/2
				if nums[mid] > nums[lenn-1] {
					left = mid + 1
				} else {
					tmp = mid - 1
				}
			}
			left = tmp + 1
		}
		if target >= nums[0] {
			// left block, find right index
			tmp := left
			for right >= tmp {
				mid := tmp + (right-tmp)/2
				if nums[mid] < nums[0] {
					right = mid - 1
				} else {
					tmp = mid + 1
				}
			}
			right = tmp - 1
		}
	}

	for right >= left {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
