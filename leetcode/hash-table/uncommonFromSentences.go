/* https://leetcode.com/problems/uncommon-words-from-two-sentences/
We are given two sentences A and B.  (A sentence is a string of space separated words.  Each word consists only of lowercase letters.)
A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.
Return a list of all uncommon words.
You may return the list in any order.

Example 1:

	Input: A = "this apple is sweet", B = "this apple is sour"
	Output: ["sweet","sour"]

Example 2:

	Input: A = "apple apple", B = "banana"
	Output: ["banana"]

Note:

	0 <= A.length <= 200
	0 <= B.length <= 200
	A and B both contain only spaces and lowercase letters.
*/

package lht

import "strings"

func uncommonFromSentences(A string, B string) []string {
	maps := map[string]int{}
	for _, words := range [][]string{strings.Split(A, " "), strings.Split(B, " ")} {
		for _, word := range words {
			if _, ok := maps[word]; ok {
				maps[word]++
			} else {
				maps[word] = 1
			}
		}
	}

	res := []string{}
	for k, v := range maps {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
