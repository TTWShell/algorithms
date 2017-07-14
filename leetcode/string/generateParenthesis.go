/* https://leetcode.com/problems/generate-parentheses/#/description
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package leetcode

func generateParenthesis(n int) []string {
	var backtrack func(res *[]string, runes []rune, open, close, max int)
	backtrack = func(res *[]string, runes []rune, open, close, max int) {
		if len(runes) == max*2 {
			*res = append(*res, string(runes))
			return
		}
		if open < max {
			backtrack(res, append(runes, '('), open+1, close, max)
		}
		if close < open {
			backtrack(res, append(runes, ')'), open, close+1, max)
		}
	}

	res := []string{}
	backtrack(&res, []rune{}, 0, 0, n)
	return res
}
