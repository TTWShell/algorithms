/* https://leetcode.com/problems/maximum-swap/description/
Given a non-negative integer, you could swap two digits at most once to get the maximum valued number. Return the maximum valued number you could get.

Example 1:
    Input: 2736
    Output: 7236
    Explanation: Swap the number 2 and the number 7.
Example 2:
    Input: 9973
    Output: 9973
    Explanation: No swap.
Note:
    The given number is in the range [0, 10^8]
*/

package lmath

import (
	"bytes"
	"strconv"
)

func maximumSwap(num int) int {
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	buckets := make([]int, 10, 10)
	for i, num := range nums {
		buckets[num] = i
	}

	nums2Int := func(nums []int) int {
		var r bytes.Buffer
		for _, num := range nums {
			r.WriteRune(rune(num) + '0')
		}
		num, _ := strconv.Atoi(r.String())
		return num
	}

	for i := 0; i < len(nums); i++ {
		for j := 9; j > nums[i]; j-- {
			if tmp := buckets[j]; tmp > i {
				nums[i], nums[tmp] = nums[tmp], nums[i]
				return nums2Int(nums)
			}
		}
	}
	return nums2Int(nums)
}
