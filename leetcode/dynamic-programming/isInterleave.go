/* https://leetcode.com/problems/interleaving-string/
Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

Example 1:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
*/

package ldp

func isInterleave(s1 string, s2 string, s3 string) bool {
	length1, length2, length3 := len(s1), len(s2), len(s3)
	if length1+length2 != length3 {
		return false
	}
	if length1 == 0 && length2 == 0 {
		return true
	}

	dp := make([][]bool, length1+1, length1+1)
	for i := range dp {
		dp[i] = make([]bool, length2+1, length2+1)
	}
	for i := 1; i <= length1; i++ {
		if s1[:i] == s3[:i] {
			dp[i][0] = true
		}
	}
	for j := 1; j <= length2; j++ {
		if s2[:j] == s3[:j] {
			dp[0][j] = true
		}
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if dp[i][j-1] == true && s3[i+j-1] == s2[j-1] {
				dp[i][j] = true
			}
			if dp[i-1][j] == true && s3[i+j-1] == s1[i-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[length1][length2]
}
