/* https://leetcode.com/problems/detect-capital/#/description
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

    All letters in this word are capitals, like "USA".
    All letters in this word are not capitals, like "leetcode".
    Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.

Example 1:
    Input: "USA"
    Output: True

Example 2:
    Input: "FlaG"
    Output: False

Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.
*/

package lstring

func detectCapitalUse(word string) bool {
	capitalCount, lowercaseCount := 0, 0
	for _, w := range word {
		if w-'A' <= 26 {
			capitalCount++
		} else if 0 <= w-'a' && w-'a' <= 26 {
			lowercaseCount++
		}
		if capitalCount > 1 && lowercaseCount > 0 {
			return false
		}
	}
	return capitalCount == len(word) ||
		lowercaseCount == len(word) ||
		(capitalCount == 1 && lowercaseCount == len(word)-1 && word[0]-'A' <= 26)
}
