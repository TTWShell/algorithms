/* https://leetcode.com/problems/sum-of-square-numbers/#/description
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
    Input: 5
    Output: True
    Explanation: 1 * 1 + 2 * 2 = 5

Example 2:
    Input: 3
    Output: False
*/

package leetcode

func judgeSquareSum(c int) bool {
	binarySearch := func(c int) int {
		var (
			mid        int
			start, end = 0, c/2 + 1
		)
		for start <= end {
			mid = start + (end-start)/2
			if temp := mid * mid; temp == c {
				break
			} else if temp > c {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		return mid
	}

	left, right := 0, binarySearch(c)
	for left <= right {
		if temp := left*left + right*right; temp == c {
			return true
		} else if temp > c {
			right--
		} else {
			left++
		}
	}
	return false
}
