/* https://leetcode.com/problems/maximum-average-subarray-ii/#/description
Given an array consisting of n integers, find the contiguous subarray whose length is greater than or equal to k that has the maximum average value.
And you need to output the maximum average value.

Example 1:
    Input: [1,12,-5,-6,50,3], k = 4
    Output: 12.75
    Explanation:
        when length is 5, maximum average value is 10.8,
        when length is 6, maximum average value is 9.16667.
        Thus return 12.75.
Note:
    1 <= k <= n <= 10,000.
    Elements of the given array will be in range [-10,000, 10,000].
    The answer with the calculation error less than 10-5 will be accepted.
*/

package leetcode

func findMaxAverage2(nums []int, k int) float64 {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	maxAverage := func(nums []int, k int) float64 {
		max := nums[k-1]
		for i := k; i < len(nums); i++ {
			if tmp := nums[i] - nums[i-k]; tmp > max {
				max = tmp
			}
		}
		return float64(max) / float64(k)
	}

	res := maxAverage(nums, k)
	for i := k + 1; i <= len(nums); i++ {
		if tmp := maxAverage(nums, i); tmp > res {
			res = tmp
		}
	}
	return res
}
