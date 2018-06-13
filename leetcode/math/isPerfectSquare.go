/* https://leetcode.com/problems/valid-perfect-square/#/description
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:
    Input: 16
    Returns: True

Example 2:
    Input: 14
    Returns: False

1 = 1
4 = 1 + 3
9 = 1 + 3 + 5
16 = 1 + 3 + 5 + 7
25 = 1 + 3 + 5 + 7 + 9
36 = 1 + 3 + 5 + 7 + 9 + 11
....
1+3+...+(2n-1) = (2n-1 + 1)n/2 = n*n

https://discuss.leetcode.com/topic/49339/o-1-time-c-solution-inspired-by-q_rsqrt
https://en.wikipedia.org/wiki/Fast_inverse_square_root
*/

package lmath

func isPerfectSquare(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}
