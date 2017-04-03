/* https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var (
		max_length int
		total      = len(s)
	)

	for i := 0; i < total; i++ {
		subString := make(map[byte]int)
		subString[s[i]] = i
		for j := i + 1; j < total; j++ {
			if _, ok := subString[s[j]]; ok {
				break
			} else {
				subString[s[j]] = j
			}
		}
		subLen := len(subString)
		if subLen > max_length {
			max_length = subLen
		}
	}
	return max_length
}

func main() {
	str := []string{"a", "au", "abcabcbb", "bbbbb", "pwwkew"}
	for _, s := range str {
		fmt.Println(s)
		fmt.Println(lengthOfLongestSubstring(s))
	}
}
