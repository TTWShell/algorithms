/* https://leetcode.com/problems/longest-word-in-dictionary/description/
Given a list of strings words representing an English Dictionary, find the longest word in words that can be built one character at a time by other words in words.
If there is more than one possible answer, return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.
Example 1:
    Input:
    words = ["w","wo","wor","worl", "world"]
    Output: "world"
    Explanation:
    The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
Example 2:
    Input:
    words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
    Output: "apple"
    Explanation:
    Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
Note:

    All the strings in the input will only contain lowercase letters.
    The length of words will be in the range [1, 1000].
    The length of words[i] will be in the range [1, 30].
*/

package lht

import "sort"

type Strings []string

func (this Strings) Len() int {
	return len(this)
}

func (this Strings) Less(i, j int) bool {
	return len(this[i]) < len(this[j]) || (len(this[i]) == len(this[j]) && this[i] > this[j])
}

func (this Strings) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func longestWord(words []string) string {
	maps := make(map[string]bool, len(words))
	for _, word := range words {
		maps[word] = true
	}

	isValid := func(word string) bool {
		for len(word) > 0 {
			if _, ok := maps[word]; !ok {
				return false
			}
			word = word[:len(word)-1]
		}
		return true
	}

	sort.Sort(Strings(words))
	longest, longestWords := len(words[len(words)-1]), []string{}
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		if !isValid(word) {
			continue
		}
		if length := len(word); length < longest {
			if len(longestWords) == 0 {
				longest, longestWords = length, []string{word}
			} else {
				break
			}
		} else if length == longest {
			longestWords = append(longestWords, word)
		}
	}
	sort.Strings(longestWords)
	return longestWords[0]
}
