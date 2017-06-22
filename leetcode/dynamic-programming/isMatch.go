/* https://leetcode.com/problems/regular-expression-matching/#/description
Implement regular expression matching with support for '.' and '*'.

    '.' Matches any single character.
    '*' Matches zero or more of the preceding element.

    The matching should cover the entire input string (not partial).

    The function prototype should be:
    bool isMatch(const char *s, const char *p)

    Some examples:
    isMatch("aa","a") → false
    isMatch("aa","aa") → true
    isMatch("aaa","aa") → false
    isMatch("aa", "a*") → true
    isMatch("aa", ".*") → true
    isMatch("ab", ".*") → true
    isMatch("aab", "c*a*b") → true
*/

package leetcode

func isMatch(s string, p string) bool {
	// 递归实现
	lens, lenp := len(s), len(p)
	if lenp == 0 {
		return lens == 0
	}

	if lenp <= 1 || (lenp > 1 && p[1] != '*') {
		return lens != 0 && (p[0] == s[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	slice := 2
	if len(p) < 2 {
		slice = 1
	}

	for len(s) != 0 && (p[0] == s[0] || p[0] == '.') {
		if isMatch(s, p[slice:]) {
			return true
		}
		s = s[1:]
	}

	return isMatch(s, p[slice:])
}
