/* https://leetcode.com/problems/longest-valid-parentheses/#/description
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

For "(()", the longest valid parentheses substring is "()", which has length = 2.

Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.

https://leetcode.com/articles/longest-valid-parentheses/
*/

package lstring

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	dp := make([]int, len(s), len(s))

	res := 0

	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}
		if s[i-1] == '(' {
			if i > 1 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
			if i-dp[i-1]-2 >= 0 {
				dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
			} else {
				dp[i] = dp[i-1] + 2
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
