/* https://leetcode.com/problems/perfect-squares/description/
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

For example, given n = 12, return 3 because 12 = 4 + 4 + 4; given n = 13, return 2 because 13 = 4 + 9.
*/

package leetcode

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
