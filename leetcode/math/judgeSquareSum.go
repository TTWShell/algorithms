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

package lmath

func judgeSquareSum(c int) bool {
	// https://en.wikipedia.org/wiki/Fermat%27s_theorem_on_sums_of_two_squares
	// https://www.zhihu.com/question/40263900
	// 自然数n可以表为两个平方数之和等价于n的每个形如p=4m+3的素因子的次数为偶数
	for i := 2; i*i <= c; i++ {
		if c%i == 0 {
			count := 0
			for c%i == 0 {
				count++
				c /= i
			}
			if i%4 == 3 && count%2 != 0 {
				return false
			}
		}
	}
	return c%4 != 3
}

/*
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
*/
