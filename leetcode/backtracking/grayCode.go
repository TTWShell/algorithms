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

import "fmt"

// https://zh.wikipedia.org/wiki/%E6%A0%BC%E9%9B%B7%E7%A0%81
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	pow := func(x, n int) int {
		var helper func(x, n int) int
		helper = func(x, n int) int {
			if n == 0 {
				return 1
			}
			if n%2 == 1 {
				return helper(x*x, n/2) * x
			} else {
				return helper(x*x, n/2)
			}
		}
		if n < 0 {
			return 1 / helper(x, -n)
		}
		return helper(x, n)
	}

	codes, index := make([][]int, pow(2, n)), 2
	codes[0], codes[1] = []int{0}, []int{1}
	for i := 2; i <= n; i++ {
		for idx, j := pow(2, i-1), index-1; j >= 0; idx, j = idx+1, j-1 {
			tmp := []int{1}
			if i-len(codes[j]) > len(tmp) {
				tmp = append(tmp, 0)
			}
			codes[idx] = append(tmp, codes[j]...)
			index++
		}
	}

	for _, sub := range codes {
		fmt.Println(n, ">>>", sub)
	}
	return []int{}
}
