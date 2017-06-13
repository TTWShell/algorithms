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
	helper := func(s string, start, end int) (frontIndex, endIndex int) {
		if start < -1 || end > len(s)-1 {
			return (start + end) / 2, (start + end) / 2
		}
		for start >= 0 && end < len(s) && s[start] == s[end] {
			start--
			end++
		}
		return start + 1, end - 1
	}

	rStartIndex, longLen := 0, 1
	for index := 0; len(s)-index >= longLen/2; index++ {
		// 按index找回文子串，以自己为中心，基数个回文串。例如s = 'aba', index = 1, return 'aba'
		if frontIndex, endIndex := helper(s, index-1, index+1); endIndex-frontIndex+1 > longLen {
			rStartIndex = frontIndex
			longLen = endIndex - frontIndex + 1

		}
		// 按index找回文字串，以自己右侧为中心，偶数个。例如：s = 'abbc', index = 1, return 'bb'
		if frontIndex, endIndex := helper(s, index, index+1); endIndex-frontIndex+1 > longLen {
			rStartIndex = frontIndex
			longLen = endIndex - frontIndex + 1
		}
	}
	return s[rStartIndex : rStartIndex+longLen]
}
