/* https://leetcode.com/problems/palindrome-partitioning-ii/
Given a string s, partition s such that every substring of
the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.

Example:

	Input: "aab"
	Output: 1
	Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/

package ldp

func minCut(s string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	length := len(s)
	dp := make([][]bool, length) // i~j是否回文
	resDp := make([]int, length, length)
	for i := 0; i < length; i++ {
		resDp[i] = i
		dp[i] = make([]bool, length)
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[j+1][i-1]) {
				dp[j][i] = true
				if j == 0 { // 0~i为回文，不需要分割
					resDp[i] = 0
				} else {
					resDp[i] = min(resDp[i], resDp[j-1]+1)
				}
			}
		}
	}
	return resDp[length-1]
}
