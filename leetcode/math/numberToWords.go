/* https://leetcode.com/problems/integer-to-english-words/
Convert a non-negative integer to its english words representation.
Given input is guaranteed to be less than 231 - 1.

Example 1:

Input: 123
Output: "One Hundred Twenty Three"
Example 2:

Input: 12345
Output: "Twelve Thousand Three Hundred Forty Five"
Example 3:

Input: 1234567
Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
Example 4:

Input: 1234567891
Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven
 Thousand Eight Hundred Ninety One"
*/

package lmath

import (
	"fmt"
	"strings"
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	var digitList = []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	var twoDigitList = []string{
		"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
		"Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	var tenList = []string{
		"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	sectionToWords := func(num int) string {
		res := ""
		if num >= 100 {
			digit := num / 100
			num = num % 100
			res = fmt.Sprintf("%s Hundred", digitList[digit])
		}
		if num >= 20 {
			digit := num / 10
			num = num % 10
			res = fmt.Sprintf("%s %s", res, tenList[digit-2])
		}
		if num >= 10 {
			res = fmt.Sprintf("%s %s", res, twoDigitList[num-10])
			num = 0
		}
		res = fmt.Sprintf("%s %s", res, digitList[num])
		return strings.TrimSpace(res)
	}

	res := ""
	for _, separator := range []string{"", "Thousand", "Million", "Billion"} {
		if num == 0 {
			break
		}
		section := num % 1000
		num = num / 1000
		if section != 0 {
			res = fmt.Sprintf("%s %s %s", sectionToWords(section), separator, res)
		}
	}
	return strings.TrimSpace(res)
}
