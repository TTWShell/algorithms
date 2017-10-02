/* https://leetcode.com/problems/gray-code/description/
The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

For example, given n = 2, return [0,1,3,2]. Its gray code sequence is:

00 - 0
01 - 1
11 - 3
10 - 2
Note:
For a given n, a gray code sequence is not uniquely defined.

For example, [0,2,3,1] is also a valid gray code sequence according to the above definition.

For now, the judge is able to judge based on one instance of gray code sequence. Sorry about that.
*/

package leetcode

import "math"

// https://zh.wikipedia.org/wiki/%E6%A0%BC%E9%9B%B7%E7%A0%81
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	codes, tmp := make([]int, int(math.Pow(2.0, float64(n)))), 2
	codes[0], codes[1] = 0, 1
	for i := 2; i <= n; i++ {
		for idx, j := tmp, tmp-1; j >= 0; idx, j = idx+1, j-1 {
			codes[idx] = tmp + codes[j]
		}
		tmp *= 2
	}
	return codes
}
