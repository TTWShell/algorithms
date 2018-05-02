/* https://leetcode.com/problems/goat-latin/description/
A sentence S is given, composed of words separated by spaces. Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)

The rules of Goat Latin are as follows:

    If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
    For example, the word 'apple' becomes 'applema'.

    If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
    For example, the word "goat" becomes "oatgma".

    Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
    For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.

Return the final sentence representing the conversion from S to Goat Latin.

Example 1:

    Input: "I speak Goat Latin"
    Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

Example 2:

    Input: "The quick brown fox jumped over the lazy dog"
    Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"

Notes:

    1. S contains only uppercase, lowercase and spaces. Exactly one space between each word.
    2. 1 <= S.length <= 150.
*/

package leetcode

import "bytes"

func toGoatLatin(S string) string {
	var res bytes.Buffer
	tmp, idx, isHead := []byte{}, 0, true
	for _, letter := range []byte(S) {
		if letter == ' ' {
			res.Write(tmp)
			idx, isHead, tmp = idx+1, true, []byte{}
			for i := idx; i > 0; i-- {
				res.WriteByte('a')
			}
			res.WriteByte(' ')
			continue
		}
		if isHead {
			isHead = false
			switch letter {
			case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
				res.WriteByte(letter)
				tmp = append(tmp, []byte("ma")...)
				continue
			default:
				tmp = append(tmp, letter)
				tmp = append(tmp, []byte("ma")...)
			}
		} else {
			res.WriteByte(letter)
		}
	}

	// tail word
	res.Write(tmp)
	for i := idx; i >= 0; i-- {
		res.WriteByte('a')
	}
	return res.String()
}
