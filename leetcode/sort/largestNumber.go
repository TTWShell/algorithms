/* https://leetcode.com/problems/largest-number/description/
Given a list of non negative integers, arrange them such that they form the largest number.

For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.

Note: The result may be very large, so you need to return a string instead of an integer.
*/

package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

type largestNumberString []string

func (s largestNumberString) Less(i, j int) bool {
	num1, _ := strconv.Atoi(s[i] + s[j])
	num2, _ := strconv.Atoi(s[j] + s[i])
	// reversed
	return num1 > num2
}

func (s largestNumberString) Len() int {
	return len(s)
}

func (s largestNumberString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	sort.Sort(largestNumberString(strNums))
	res := strings.Join(strNums, "")

	idx := 0
	for idx < len(res)-1 && res[idx] == '0' {
		idx++
	}
	return res[idx:]
}
