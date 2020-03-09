/* https://leetcode.com/problems/distinct-subsequences/
Given a string S and a string T, count the number of distinct subsequences
of S which equals T.

A subsequence of a string is a new string which is formed from
the original string by deleting some (can be none) of
the characters without disturbing the relative positions of
the remaining characters. (ie, "ACE" is a subsequence of "ABCDE"
while "AEC" is not).

Example 1:

	Input: S = "rabbbit", T = "rabbit"
	Output: 3
	Explanation:

	As shown below, there are 3 ways you can generate "rabbit" from S.
	(The caret symbol ^ means the chosen letters)

	rabbbit
	^^^^ ^^
	rabbbit
	^^ ^^^^
	rabbbit
	^^^ ^^^

Example 2:

	Input: S = "babgbag", T = "bag"
	Output: 5
	Explanation:

	As shown below, there are 5 ways you can generate "bag" from S.
	(The caret symbol ^ means the chosen letters)

	babgbag
	^^ ^
	babgbag
	^^    ^
	babgbag
	^    ^^
	babgbag
	^  ^^
	babgbag
		^^^
*/

package ldp

func numDistinct(s string, t string) int {
	row, col := len(t)+1, len(s)+1
	dp := make([][]int, row, row)
	for i := range dp {
		dp[i] = make([]int, col, col)
	}
	for j := 0; j < col; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			agoRow, agoCol := i-1, j-1
			if s[agoCol] == t[agoRow] {
				dp[i][j] = dp[i][agoCol] + dp[agoRow][agoCol]
			} else {
				dp[i][j] = dp[i][agoCol]
			}
		}
	}
	return dp[row-1][col-1]
}
