/* https://leetcode.com/problems/verifying-an-alien-dictionary/
In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order.
The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet,
return true if and only if the given words are sorted lexicographicaly in this alien language.


Example 1:

	Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
	Output: true
	Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.

Example 2:

	Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
	Output: false
	Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.

Example 3:

	Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
	Output: false
	Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app",
	because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character (More info).

Note:

	1 <= words.length <= 100
	1 <= words[i].length <= 20
	order.length == 26
	All characters in words[i] and order are english lowercase letters.
*/

package lht

func isAlienSorted(words []string, order string) bool {
	orderMap := map[byte]int{}
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		pre, next, flag := words[i-1], words[i], false
		for i, j := 0, 0; i < len(pre) && j < len(next); i, j = i+1, j+1 {
			if ip, in := orderMap[pre[i]], orderMap[next[j]]; ip > in {
				return false
			} else if ip < in {
				flag = true
				break
			}
		}
		if flag || len(pre) <= len(next) {
			continue
		}
		return false
	}
	return true
}
