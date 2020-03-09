/* https://leetcode.com/problems/word-break-ii/
Given a non-empty string s and a dictionary wordDict containing a list
of non-empty words, add spaces in s to construct a sentence where
each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:

	Input:
	s = "catsanddog"
	wordDict = ["cat", "cats", "and", "sand", "dog"]
	Output:
	[
	"cats and dog",
	"cat sand dog"
	]

Example 2:

	Input:
	s = "pineapplepenapple"
	wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
	Output:
	[
	"pine apple pen apple",
	"pineapple pen apple",
	"pine applepen apple"
	]
	Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:

	Input:
	s = "catsandog"
	wordDict = ["cats", "dog", "sand", "and", "cat"]
	Output:
	[]
*/

package lbacktracking

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
	}

	cache := map[string][][]string{}

	var helper func(startIdx int, s string) [][]string
	helper = func(startIdx int, s string) [][]string {
		if value, ok := cache[s]; ok {
			return value
		}

		res := [][]string{}
		if _, ok := dict[s]; ok {
			res = append(res, []string{s})
		}

		for i := 1; i < len(s); i++ {
			word := s[:i]
			if _, ok := dict[word]; ok {
				subs := helper(startIdx+i, s[i:])
				for _, sub := range subs {
					res = append(res, append([]string{word}, sub...))
				}
			}
		}
		cache[s] = res
		return res
	}

	words := helper(0, s)

	l := len(words)
	if l == 0 {
		return []string{}
	}
	res := make([]string, l, l)
	for i := range words {
		res[i] = strings.Join(words[i], " ")
	}
	return res
}
