/* https://leetcode.com/problems/factorial-trailing-zeroes/#/description
Given an integer n, return the number of trailing zeroes in n!.

Note: Your solution should be in logarithmic time complexity.

给定一个整数n，返回n!（n的阶乘）数字中的后缀0的个数。

注意：你的解法应该满足多项式时间复杂度。

思路：
	欲求数字后缀中0的个数，本质就是求n * (n-1) ... 1中10的位数。
	其中10 ＝ 5 ＊ 2，显而易见，2的个数 >= 5的个数 恒成立。所以转化为：
	n!后缀0的个数 = n!质因子中5的个数
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
