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

import "math"

func findMaxAverage2(nums []int, k int) float64 {
	check := func(nums []int, val float64, k int) bool {
		sum := 0.0
		for i := 0; i < k; i++ {
			sum += float64(nums[i]) - val
		}
		if sum >= 0 {
			return true
		}

		prev, min := 0.0, 0.0
		for i := k; i < len(nums); i++ {
			sum += float64(nums[i]) - val
			prev += float64(nums[i-k]) - val
			min = math.Min(prev, min)
			if sum >= min {
				return true
			}
		}
		return false
	}

	min, max := math.MaxFloat64, float64(math.MinInt32)
	for _, x := range nums {
		num := float64(x)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	prevMid, err := max, 1.0
	for err > 0.00001 {
		mid := min + 0.5*(max-min)
		if check(nums, mid, k) {
			min = mid
		} else {
			max = mid
		}
		prevMid, err = mid, math.Abs(prevMid-mid)
	}

	return min
}

/*
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
*/
