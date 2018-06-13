/* https://leetcode.com/problems/ugly-number-ii/description/
Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

Note that 1 is typically treated as an ugly number, and n does not exceed 1690.
*/

package lmath

func nthUglyNumber(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := make([]int, n, n)
	res[0] = 1
	i2, i3, i5 := 0, 0, 0
	for idx := 1; idx < n; idx++ {
		m2, m3, m5 := res[i2]*2, res[i3]*3, res[i5]*5
		mn := min(m2, min(m3, m5))
		if mn == m2 {
			i2++
		}
		if mn == m3 {
			i3++
		}
		if mn == m5 {
			i5++
		}
		res[idx] = mn
	}
	return res[n-1]
}
