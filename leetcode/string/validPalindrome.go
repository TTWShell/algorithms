/* https://leetcode.com/problems/valid-palindrome-ii/description/
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
    Input: "aba"
    Output: True

Example 2:
    Input: "abca"
    Output: True
    Explanation: You could delete the character 'c'.

Note:
    The string will only contain lowercase characters a-z. The maximum length of the string is 50000.
*/

package lstring

func validPalindrome(s string) bool {
	var isPalindrome func(s string, start, end, count int) bool
	isPalindrome = func(s string, start, end, count int) bool {
		for start < end {
			if s[start] != s[end] {
				if count == 0 {
					return false
				}
				count--
				return isPalindrome(s, start+1, end, count) || isPalindrome(s, start, end-1, count)
			}
			start++
			end--
		}
		return true
	}
	return isPalindrome(s, 0, len(s)-1, 1)
}
