/* https://leetcode.com/problems/increasing-triplet-subsequence/
Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Formally the function should:

Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.

Example 1:

	Input: [1,2,3,4,5]
	Output: true

Example 2:

	Input: [5,4,3,2,1]
	Output: false
*/

package larray

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := nums[0]
	minMid := math.MaxInt32
	for i := 1; i < len(nums); i++ {
		if nums[i] > minMid {
			return true
		}
		if min > nums[i-1] {
			min = nums[i-1]
		}
		if nums[i] > min {
			if nums[i] < minMid {
				minMid = nums[i]
			}
		}
	}
	return false
}
