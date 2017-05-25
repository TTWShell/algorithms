/* https://leetcode.com/problems/factorial-trailing-zeroes/#/description
Given an integer n, return the number of trailing zeroes in n!.

Note: Your solution should be in logarithmic time complexity.

给定一个整数n，返回n!（n的阶乘）数字中的后缀0的个数。

注意：你的解法应该满足多项式时间复杂度。
*/

package leetcode

func trailingZeroes(n int) int {
	r := 0
	for n > 0 {
		n /= 5
		r += n
	}
	return r
}
