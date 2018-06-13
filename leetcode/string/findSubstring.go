/* https://leetcode.com/problems/substring-with-concatenation-of-all-words/#/description
You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).
*/

package lstring

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(s) == 0 || len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return res
	}

	initwMaps := func(wMaps map[string]int, words []string) {
		for _, word := range words {
			wMaps[word] = 0
		}
		for _, word := range words {
			wMaps[word]++
		}
	}

	wMaps := make(map[string]int, len(words))
	lens, lenws, lenw := len(s), len(words), len(words[0])

	for slide := 0; slide < lenw; slide++ {
		initwMaps(wMaps, words)
		i, count := 0, 0
		for i <= lens-slide-lenws*lenw {
			tmp := s[i+slide+count*lenw : i+slide+(count+1)*lenw]
			if num, ok := wMaps[tmp]; ok && num != 0 {
				wMaps[tmp]--
				count++
				if count == lenws {
					res = append(res, i+slide)
					wMaps[s[i+slide:i+slide+lenw]]++
					count--
					i += lenw
				}
			} else if ok {
				wMaps[s[i+slide:i+slide+lenw]]++
				count--
				i += lenw

			} else {
				if count != 0 {
					initwMaps(wMaps, words)
				}
				i += lenw * (count + 1)
				count = 0
			}
		}
	}

	return res
}
