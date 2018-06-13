/* https://leetcode.com/problems/summary-ranges/description/
Given a sorted integer array without duplicates, return the summary of its ranges.

Example 1:

	Input: [0,1,2,4,5,7]
	Output: ["0->2","4->5","7"]

Example 2:

	Input: [0,2,3,4,6,8,9]
	Output: ["0","2->4","6","8->9"]
*/

package larray

import "strconv"

func summaryRanges(nums []int) []string {
	length := len(nums)
	if length == 0 {
		return []string{}
	}

	res := []string{}
	for left, i := nums[0], 1; i <= len(nums); i++ {
		if i == length || nums[i] != nums[i-1]+1 {
			right := nums[i-1]
			if left == right {
				res = append(res, strconv.Itoa(left))
			} else {
				res = append(res, strconv.Itoa(left)+"->"+strconv.Itoa(right))
			}
			if i < length {
				left = nums[i]
			}
		}
	}
	return res
}
