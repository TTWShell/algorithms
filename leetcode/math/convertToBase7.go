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

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	r := ""
	for num > 0 {
		r = strconv.Itoa(num%7) + r
		num /= 7
	}
	if isNegative == true {
		return "-" + r
	}
	return r
}
