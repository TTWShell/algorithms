/* https://leetcode.com/problems/reverse-integer/#/description
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

Note:
The input is assumed to be a 32-bit signed integer. Your function should return 0 when the reversed integer overflows.
*/
package leetcode

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result < -2147483648 || result > 2147483647 {
		return 0
	}
	return result
}
