/* https://leetcode.com/problems/unique-morse-code-words/description/
International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter.
For example, "cab" can be written as "-.-.-....-", (which is the concatenation "-.-." + "-..." + ".-").
We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.

Example:

	Input: words = ["gin", "zen", "gig", "msg"]
	Output: 2
	Explanation:
	The transformation of each word is:
		"gin" -> "--...-."
		"zen" -> "--...-."
		"gig" -> "--...--."
		"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".


Note:

    1. The length of words will be at most 100.
    2. Each words[i] will have length in range [1, 12].
    3. words[i] will only consist of lowercase letters.
*/

package leetcode

import "bytes"

func uniqueMorseRepresentations(words []string) int {
	table := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-",
		"..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	dict := make(map[string]int)

	for _, word := range words {
		tmp := bytes.Buffer{}
		for _, char := range word {
			tmp.WriteString(table[char-'a'])
		}
		dict[tmp.String()]++
	}
	return len(dict)
}
