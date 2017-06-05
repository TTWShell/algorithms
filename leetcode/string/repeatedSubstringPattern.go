/* https://leetcode.com/problems/repeated-substring-pattern/#/description
Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.
You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.

Example 1:
    Input: "abab"
    Output: True

    Explanation: It's the substring "ab" twice.

Example 2:
    Input: "aba"
    Output: False

Example 3:
    Input: "abcabcabcabc"
    Output: True

    Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
*/

package leetcode

func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	}

	subLen := 1
	for i := subLen; i < len(s) && subLen <= len(s)/2; {
		if len(s)%subLen != 0 {
			subLen += 1
			i = subLen
			continue
		}

		if string(s[i:i+subLen]) == string(s[0:subLen]) {
			i += subLen
		} else {
			subLen += 1
			i = subLen
		}
	}

	if subLen <= len(s)/2 {
		return true
	}
	return false
}
