/* https://leetcode.com/problems/word-pattern/#/description
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Examples:
pattern = "abba", str = "dog cat cat dog" should return true.
pattern = "abba", str = "dog cat cat fish" should return false.
pattern = "aaaa", str = "dog cat cat dog" should return false.
pattern = "abba", str = "dog dog dog dog" should return false.
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.
*/

package lht

import "strings"

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(pattern) != len(words) {
		return false
	}

	pmap := make(map[rune]string)
	smap := make(map[string]rune)
	for i, p := range pattern {
		if v, ok := pmap[p]; !ok {
			pmap[p] = words[i]
		} else if v != words[i] {
			return false
		}
		if v, ok := smap[words[i]]; !ok {
			smap[words[i]] = p
		} else if v != p {
			return false
		}
	}

	return true
}
