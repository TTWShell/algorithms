/* https://leetcode.com/problems/longest-palindromic-substring/#/description
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
    Input: "babad"
    Output: "bab"
    Note: "aba" is also a valid answer.

Example:
    Input: "cbbd"
    Output: "bb"
*/

package leetcode

func longestPalindrome(s string) string {
	helper := func(s string, start, end int) string {
		if start < -1 || end > len(s)-1 {
			return string(s[(start+end)/2])
		}
		for start >= 0 && end < len(s) {
			if s[start] != s[end] {
				break
			}
			start--
			end++
		}
		return s[start+1 : end]
	}

	result := ""
	for index := 0; index < len(s); index++ {
		// 按index找回文子串，以自己为中心，基数个回文串。例如s = 'aba', index = 1, return 'aba'
		if temp := helper(s, index-1, index+1); len(temp) > len(result) {
			result = temp
		}
		// 按index找回文字串，以自己右侧为中心，偶数个。例如：s = 'abbc', index = 1, return 'bb'
		if temp := helper(s, index, index+1); len(temp) > len(result) {
			result = temp
		}
	}
	return result
}
