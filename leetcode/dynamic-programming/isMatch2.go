/* https://leetcode.com/problems/wildcard-matching/description/
Implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
    isMatch("aa","a") → false
    isMatch("aa","aa") → true
    isMatch("aaa","aa") → false
    isMatch("aa", "*") → true
    isMatch("aa", "a*") → true
    isMatch("ab", "?*") → true
    isMatch("aab", "c*a*b") → false
*/

package leetcode

func isMatch2(s string, p string) bool {
	si, pi, match, stari := 0, 0, 0, -1
	for si < len(s) {
		if pi < len(p) && (p[pi] == '?' || s[si] == p[pi]) {
			si++
			pi++
		} else if pi < len(p) && p[pi] == '*' {
			stari = pi
			match = si
			pi++
		} else if stari != -1 {
			pi = stari + 1
			match++
			si = match
		} else {
			return false
		}
	}
	for pi < len(p) && p[pi] == '*' {
		pi++
	}
	return pi == len(p)
}
