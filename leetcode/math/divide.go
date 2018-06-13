/* https://leetcode.com/problems/divide-two-integers/#/description
Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.
*/

package lmath

func divide(dividend int, divisor int) int {
	// 参考https://discuss.leetcode.com/category/37/divide-two-integers
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	sign := true
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = false
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	res := 0
	for dividend >= divisor {
		temp, multiple := divisor, 1
		for (temp << 1) < dividend {
			temp = temp << 1
			multiple = multiple << 1
		}
		dividend = dividend - temp
		res += multiple
	}

	if !sign {
		return -res
	}
	return res
}
