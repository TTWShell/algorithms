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

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	sort.Slice(strNums, func(i, j int) bool {
		// reversed
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	res := strings.Join(strNums, "")
	if res[0] == '0' {
		return "0"
	}
	return res
}
