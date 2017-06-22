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
		// 如果下一个p不为*: 则检查当前s与p是否匹配，匹配的话继续处理isMatch(s[1:], p[1:])
		return lens != 0 && (p[0] == s[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	// panic: runtime error: slice bounds out of range
	slice := 2
	if len(p) < 2 {
		slice = 1
	}

	// 下一个p为*，则循环处理s
	for len(s) != 0 && (p[0] == s[0] || p[0] == '.') {
		if isMatch(s, p[slice:]) {
			// 当前s与下二个p也都匹配，isMatch(s, p[2:])
			return true
		}
		s = s[1:]
	}

	// isMatch(s, p[2:])
	return isMatch(s, p[slice:])
}
