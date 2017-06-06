/* https://leetcode.com/problems/max-consecutive-ones/#/description
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
    Input: [1,1,0,1,1,1]
    Output: 3
    Explanation: The first two digits or the last three digits are consecutive 1s.
        The maximum number of consecutive 1s is 3.

Note:

    The input array will only contain 0 and 1.
    The length of input array is a positive integer and will not exceed 10,000
*/

package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	r, cur := 0, 0
	for _, num := range nums {
		if num == 1 {
			cur++
		} else {
			r = max(r, cur)
			cur = 0
		}
	}
	r = max(r, cur)
	return r
}
