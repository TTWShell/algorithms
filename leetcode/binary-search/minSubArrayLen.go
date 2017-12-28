/* https://leetcode.com/problems/minimum-size-subarray-sum/description/
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

For example, given the array [2,3,1,2,4,3] and s = 7,
the subarray [4,3] has the minimal length under the problem constraint.

More practice:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
*/

package leetcode

// Two Pointers
func minSubArrayLen(s int, nums []int) int {
	sums := make([]int, len(nums)+1)

	l, r, minLen := 0, 0, 0
	for r < len(sums) {
		if sums[r]-sums[l] >= s {
			if tmp := r - l; tmp < minLen || minLen == 0 {
				minLen = tmp
			}
			l++
		} else if r++; r < len(sums) {
			sums[r] = sums[r-1] + nums[r-1]
		}
	}
	return minLen
}
