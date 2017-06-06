/* https://leetcode.com/problems/keyboard-row/#/description
Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

https://leetcode.com/static/images/problemset/keyboard.png

Example 1:
    Input: ["Hello", "Alaska", "Dad", "Peace"]
    Output: ["Alaska", "Dad"]
Note:
    You may use one character in the keyboard more than once.
    You may assume the input string will only contain letters of alphabet.
*/

package leetcode

func findWords(words []string) []string {
	alphabets := map[rune]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}

	var dstWords []string
	for _, word := range words {
		row, isOneRow := 0, true
		for i, letter := range word {
			if 'A' <= letter && letter <= 'Z' {
				letter = letter + 'a' - 'A'
			}
			if temp := alphabets[letter]; i == 0 {
				row = temp
			} else if temp != row {
				isOneRow = false
				break
			}
		}
		if isOneRow {
			dstWords = append(dstWords, word)
		}
	}
	return dstWords
}
