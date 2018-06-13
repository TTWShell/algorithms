/* https://leetcode.com/problems/largest-palindrome-product/description/
Find the largest palindrome made from the product of two n-digit numbers.

Since the result could be very large, you should return the largest palindrome mod 1337.

Example:

Input: 2

Output: 987

Explanation: 99 x 91 = 9009, 9009 % 1337 = 987

Note:

The range of n is [1,8].
*/

package lothers

import (
	"math"
	"strconv"
)

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	reverse := func(s string) string {
		res := []byte(s)
		for i := 0; i < len(res)/2; i++ {
			res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
		}
		return string(res)
	}

	up := int64(math.Pow10(n)) - 1
	low := up/10 + 1

	for i := up; i >= low; i-- {
		prod, _ := strconv.ParseInt(strconv.FormatInt(i, 10)+reverse(strconv.FormatInt(i, 10)), 10, 64)
		for j := up; j >= prod/j; j-- {
			if prod%j == 0 {
				return int(prod % 1337)
			}
		}
	}
	return 0
}

/*
func largestPalindrome(n int) int {
	res := []int64{9, 9009, 906609, 99000099, 9966006699, 999000000999, 99956644665999, 9999000000009999}
	return int(res[n-1] % int64(1337))
}
*/
