/* https://leetcode.com/problems/palindromic-substrings/#/description
Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:
    Input: "abc"
    Output: 3
    Explanation: Three palindromic strings: "a", "b", "c".
Example 2:
    Input: "aaa"
    Output: 6
    Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
Note:
    The input string length won't exceed 1000.
*/

package leetcode

func countSubstrings(s string) int {
	isPalindromic := func(s string) bool {
		for i, j := 0, len(s)-1; i <= j; {
			if s[i] != s[len(s)-1-i] {
				return false
			}
			i++
			j--
		}
		return true
	}

	res := len(s)
	for step := 2; step <= len(s); step++ {
		for start := 0; start <= len(s)-step; start++ {
			if isPalindromic(s[start : start+step]) {
				res++
			}
		}
	}
	return res
}
