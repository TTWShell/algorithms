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

package ldp

func isMatch(s string, p string) bool {
	/* 动态规划 https://discuss.leetcode.com/topic/40371/easy-dp-java-solution-with-detailed-explanation
		dp[i][j]表示s[0,i)和p[0,j)是否match
		1. If p.charAt(j) == s.charAt(i) :  dp[i][j] = dp[i-1][j-1];
		2. If p.charAt(j) == '.' : dp[i][j] = dp[i-1][j-1];
		3. If p.charAt(j) == '*':
	   here are two sub conditions:
		   1. if p.charAt(j-1) != s.charAt(i) : dp[i][j] = dp[i][j-2]  //in this case, a* only counts as empty
		   2. if p.charAt(i-1) == s.charAt(i) or p.charAt(i-1) == '.':
						  dp[i][j] = dp[i-1][j]    //in this case, a* counts as multiple a
					   or dp[i][j] = dp[i][j-1]   // in this case, a* counts as single a
					   or dp[i][j] = dp[i][j-2]   // in this case, a* counts as empty
	*/
	lens, lenp := len(s), len(p)
	if lens == 0 && lenp == 0 {
		return true
	}

	dp := make([][]bool, lens+1)
	for i := 0; i <= lens; i++ {
		dp[i] = make([]bool, lenp+1)
	}

	dp[0][0] = true // 初始认为s、p前面都有一个空串，并认为匹配
	for i := 0; i < lenp; i++ {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < lens; i++ {
		for j := 0; j < lenp; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]
				}
			}
		}
	}

	return dp[lens][lenp]
}

/*
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
*/
