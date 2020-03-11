/* https://leetcode.com/problems/shortest-palindrome/
Given a string s, you are allowed to convert it to a palindrome by
 adding characters in front of it. Find and return the shortest
 palindrome you can find by performing this transformation.

Example 1:

	Input: "aacecaaa"
	Output: "aaacecaaa"

Example 2:

	Input: "abcd"
	Output: "dcbabcd"
*/

package lstring

func shortestPalindrome(s string) string {
	isPalindrome := func(start, end int) bool {
		for i := start; i < start+(end-start)/2; i++ {
			if s[i] != s[end-1-i] {
				return false
			}
		}
		return true
	}

	end := len(s)
	for ; end >= 1; end-- {
		if isPalindrome(0, end) {
			break
		}
	}

	tmp, idx := make([]byte, len(s)-end), 0
	for i := len(s) - 1; i >= end; i-- {
		tmp[idx] = s[i]
		idx++
	}
	return string(tmp) + s
}
