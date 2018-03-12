/* https://leetcode.com/problems/perfect-squares/description/
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

For example, given n = 12, return 3 because 12 = 4 + 4 + 4; given n = 13, return 2 because 13 = 4 + 9.
*/

package leetcode

/*
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		dp[i*i] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; i+j*j <= n; j++ {
			if dp[i+j*j] == 0 || dp[i+j*j] > dp[i]+1 {
				dp[i+j*j] = dp[i] + 1
			}
		}
	}
	return dp[n]
}
*/

import "math"

func numSquares(n int) int {
	for n%4 == 0 {
		n /= 4
	}

	if n%8 == 7 {
		return 4
	}

	for i := 0; i*i <= n; i++ {
		j := int(math.Sqrt(float64(n - i*i)))
		if i*i+j*j == n {
			res := 0
			if i > 0 {
				res++
			}
			if j > 0 {
				res++
			}
			return res
		}
	}
	return 3
}
