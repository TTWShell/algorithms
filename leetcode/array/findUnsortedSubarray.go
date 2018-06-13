/* https://leetcode.com/problems/shortest-unsorted-continuous-subarray/#/description
Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order,
then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
    Input: [2, 6, 4, 8, 10, 9, 15]
    Output: 5
    Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.

Note:
    Then length of the input array is in range [1, 10,000].
    The input array may contain duplicates, so ascending order here means <=.
*/

package larray

/*
import "sort"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	sorted := make([]int, n, n)
	copy(sorted, nums)
	sort.Ints(sorted)

	i, j := 0, n-1
	for i < n && nums[i] == sorted[i] {
		i++
	}
	for j > i && nums[j] == sorted[j] {
		j--
	}
	return j + 1 - i
}
*/

func findUnsortedSubarray(nums []int) int {
	n, left, right := len(nums), -1, -2
	max, min := nums[0], nums[n-1]

	for i := 1; i < n; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			right = i
		}
		if end := n - 1 - i; nums[end] <= min {
			min = nums[end]
		} else {
			left = end
		}
	}

	return right - left + 1
}
