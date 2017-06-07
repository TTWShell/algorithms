/* https://leetcode.com/problems/base-7/#/description
Given an integer, return its base 7 string representation.

Example 1:
    Input: 100
    Output: "202"
Example 2:
    Input: -7
    Output: "-10"
Note: The input will be in range of [-1e7, 1e7].
*/

package leetcode

import (
	"strconv"
	"strings"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	reverse := func(chars []string) {
		for i := 0; i < len(chars)/2; i++ {
			j := len(chars) - i - 1
			chars[i], chars[j] = chars[j], chars[i]
		}
	}

	isNegative := false
	if num < 0 {
		isNegative, num = true, -num
	}

	var symbols []string
	for num > 0 {
		symbols = append(symbols, strconv.Itoa(num%7))
		num /= 7
	}
	if isNegative == true {
		symbols = append(symbols, "-")
	}

	reverse(symbols)

	return strings.Join(symbols, "")
}
