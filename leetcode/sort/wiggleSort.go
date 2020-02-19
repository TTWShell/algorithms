/* https://leetcode.com/problems/wiggle-sort-ii/
Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

Example 1:

Input: nums = [1, 5, 1, 1, 6, 4]
Output: One possible answer is [1, 4, 1, 5, 1, 6].
Example 2:

Input: nums = [1, 3, 2, 2, 3, 1]
Output: One possible answer is [2, 3, 1, 3, 1, 2].
Note:
You may assume all input has valid answer.

Follow Up:
Can you do it in O(n) time and/or in-place with O(1) extra space?
*/

package lsort

import (
	"sort"
)

func wiggleSort(nums []int) {
	a := make([]int, len(nums))
	copy(a, nums)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	var head int
	for i := 1; i < len(nums); i += 2 {
		head, a = a[0], a[1:]
		nums[i] = head
	}
	for i := 0; i < len(nums); i += 2 {
		head, a = a[0], a[1:]
		nums[i] = head
	}
}
