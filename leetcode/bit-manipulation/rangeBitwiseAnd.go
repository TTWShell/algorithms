/* https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

For example, given the range [5, 7], you should return 4.
*/

package leetcode

func rangeBitwiseAnd(m int, n int) int {
	for offset := uint(0); m&n != 0; offset++ {
		if m == n {
			// shift count type int, must be unsigned integer
			return int(uint(m) << offset)
		}
		m >>= 1
		n >>= 1
	}
	return 0
}
