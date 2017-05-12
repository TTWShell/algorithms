/* https://leetcode.com/problems/length-of-last-word/#/description
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

For example,
Given s = "Hello World",
return 5.
*/

package leetcode

import "unicode"

func lengthOfLastWord(s string) int {
	len, r, f := len(s), 0, false
	for i := len - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			f = true
			r++
		} else {
			if f == true {
				break
			}
		}
	}
	return r
}
