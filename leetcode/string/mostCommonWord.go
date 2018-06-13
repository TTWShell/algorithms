/* https://leetcode.com/problems/most-common-word/description/
Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.

    Example:
    Input:
    paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
    banned = ["hit"]
    Output: "ball"
    Explanation:
    "hit" occurs 3 times, but it is a banned word.
    "ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
    Note that words in the paragraph are not case sensitive,
    that punctuation is ignored (even if adjacent to words, such as "ball,"),
    and that "hit" isn't the answer even though it occurs more because it is banned.

Note:

    1 <= paragraph.length <= 1000.
    1 <= banned.length <= 100.
    1 <= banned[i].length <= 10.
    The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
    paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
    Different words in paragraph are always separated by a space.
    There are no hyphens or hyphenated words.
    Words only consist of letters, never apostrophes or other punctuation symbols.
*/

package lstring

import "bytes"

func mostCommonWord(paragraph string, banned []string) string {
	bannedWords := make(map[string]int, len(banned))
	for _, word := range banned {
		bannedWords[word] = 1
	}
	words := make(map[string]int)

	res, resTimes := "", 0
	var tmp bytes.Buffer
	for _, letter := range append([]byte(paragraph), ' ') {
		if 65 <= letter && letter <= 90 {
			letter = letter + 32
		}
		if 97 <= letter && letter <= 122 {
			tmp.WriteByte(letter)
		} else if word := tmp.String(); len(word) != 0 {
			if _, ok := bannedWords[word]; !ok {
				words[word]++
				if words[word] > resTimes {
					res, resTimes = word, words[word]
				}
			}
			tmp = bytes.Buffer{}
		}
	}
	return res
}
