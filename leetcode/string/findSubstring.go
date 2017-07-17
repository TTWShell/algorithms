/* https://leetcode.com/problems/substring-with-concatenation-of-all-words/#/description
You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).
*/

package leetcode

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}

	wMaps := make(map[string]int, len(words))
	for _, word := range words {
		wMaps[word]++
	}

	lens, lenws, lenw := len(s), len(words), len(words[0])
	for i := 0; i <= lens-lenws*lenw; i++ {
		tmpMaps := make(map[string]int)
		count := 0
		for j := 0; j < lenws; j++ {
			tmp := s[i+j*lenw : i+j*lenw+lenw]
			tmpMaps[tmp]++
			if v, ok := wMaps[tmp]; !ok || v < tmpMaps[tmp] {
				break
			}
			count++
		}
		if count == lenws {
			res = append(res, i)
		}
	}

	return res
}
