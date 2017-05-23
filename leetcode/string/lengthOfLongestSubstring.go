/* https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/
package leetcode

func lengthOfLongestSubstring(s string) int {
	var (
		ok           bool
		i, j, max, d int
		total        = len(s)
		subString    = make(map[byte]int)
	)
	for ; j < total; j++ {
		if _, ok = subString[s[j]]; ok {
			if subString[s[j]] > i {
				i = subString[s[j]]
			}
		}
		if d = j - i + 1; d > max {
			max = d
		}
		subString[s[j]] = j + 1
	}
	return max
}
