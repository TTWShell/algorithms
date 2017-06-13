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
	rStartIndex, longLen := 0, 1
	for index := 0; len(s)-index >= longLen/2; index++ {
		frontIndex, endIndex := index, index
		for endIndex < len(s)-1 && s[endIndex] == s[endIndex+1] {
			endIndex++
		}
		for frontIndex > 0 && endIndex < len(s)-1 && s[frontIndex-1] == s[endIndex+1] {
			frontIndex--
			endIndex++
		}
		if endIndex-frontIndex+1 > longLen {
			rStartIndex = frontIndex
			longLen = endIndex - frontIndex + 1
		}
	}
	return s[rStartIndex : rStartIndex+longLen]
}
