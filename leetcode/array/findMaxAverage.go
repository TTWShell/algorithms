/* https://leetcode.com/problems/maximum-average-subarray-i/#/description
Given an array consisting of n integers, find the contiguous subarray of given length k that has the maximum average value. And you need to output the maximum average value.

Example 1:
    Input: [1,12,-5,-6,50,3], k = 4
    Output: 12.75
    Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
Note:
    1 <= k <= n <= 30,000.
    Elements of the given array will be in the range [-10,000, 10,000].
*/

package leetcode

func findMaxAverage(nums []int, k int) float64 {
	var max, curSum int
	for i := 0; i < k && i < len(nums); i++ {
		curSum += nums[i]
	}
	max = curSum

	for i := k; i < len(nums); i++ {
		curSum += nums[i] - nums[i-k]
		if curSum > max {
			max = curSum
		}
	}

	return float64(max) / float64(k)
}
