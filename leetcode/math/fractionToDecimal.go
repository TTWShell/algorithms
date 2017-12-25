/* https://leetcode.com/problems/fraction-to-recurring-decimal/description/
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

For example,

    Given numerator = 1, denominator = 2, return "0.5".
    Given numerator = 2, denominator = 1, return "2".
    Given numerator = 2, denominator = 3, return "0.(6)".
*/

package leetcode

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	res := []string{}
	if numerator*denominator < 0 {
		res = append(res, "-")
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}

	integer, fractional := numerator/denominator, numerator%denominator
	res = append(res, strconv.Itoa(integer))
	if fractional != 0 {
		res = append(res, ".")
	}

	idx := len(res)
	maps := map[int]int{fractional: idx}
	idx++

	for ; fractional != 0; idx++ {
		res = append(res, strconv.Itoa(fractional*10/denominator))
		fractional = fractional * 10 % denominator
		if i, ok := maps[fractional]; ok {
			tmp := []string{}
			tmp = append(tmp, res[:i]...)
			tmp = append(tmp, "(")
			tmp = append(tmp, res[i:]...)
			tmp = append(tmp, ")")
			res = tmp
			break
		}
		maps[fractional] = idx
	}

	return strings.Join(res, "")
}
