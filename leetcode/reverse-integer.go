/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

Note:
The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
*/
package leetcode

func reverse(x int) int {
	sign, result := 1, 0
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result < 0 || result > 2147483647 {
		return 0
	}
	return result * sign
}
