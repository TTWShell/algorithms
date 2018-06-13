/* https://leetcode.com/problems/string-to-integer-atoi/#/description
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

Requirements for atoi:
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.
*/

package lmath

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	/* 题目要求蛮多的
	1. 首先需要丢弃字符串前面的空格；
	2. 然后可能有正负号（注意只取一个，如果有多个正负号，那么说这个字符串是无法转换的，返回0。比如测试用例里就有个“+-2”）；
	3. 字符串可以包含0~9以外的字符，如果遇到非数字字符，那么只取该字符之前的部分，如“-00123a66”返回为“-123”；
	4. 如果超出int的范围，返回边界值（2147483647或-2147483648）。
	*/
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}

	neg := false
	if str[0] == '-' || str[0] == '+' {
		neg = true && str[0] != '+'
		str = str[1:]
	}

	un := uint(0)
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			// 字符串中出现数字以外的字符
			break
		}

		if un > math.MaxUint32/10 {
			// ? overflow
			if neg {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		un *= 10
		un1 := un + uint(str[i]) - '0'
		if un1 < un {
			// overflow
			return 0
		}
		un = un1
	}

	if !neg && un > math.MaxInt32 {
		return math.MaxInt32
	}
	if neg && un > math.MaxInt32+1 {
		return math.MinInt32
	}

	n := int(un)
	if neg {
		n = -n
		// 如果小于最小int, 那么就返回最小int
		if n < math.MinInt32 {
			return math.MinInt32
		}
	}

	return n
}
