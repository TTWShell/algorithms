/* https://leetcode.com/problems/sum-of-two-integers/#/description
Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

Example:
Given a = 1 and b = 2, return 3.
*/

package lbm

func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	for b != 0 {
		carry := a & b
		a, b = a^b, carry<<1
	}
	return a
}
