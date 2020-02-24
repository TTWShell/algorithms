/* https://leetcode.com/problems/reverse-words-in-a-string/
Given an input string, reverse the string word by word.

Example 1:

	Input: "the sky is blue"
	Output: "blue is sky the"

Example 2:

	Input: "  hello world!  "
	Output: "world! hello"
	Explanation: Your reversed string should not contain leading or
	trailing spaces.

Example 3:

	Input: "a good   example"
	Output: "example good a"
	Explanation: You need to reduce multiple spaces between two words
	to a single space in the reversed string.

Note:

	A word is defined as a sequence of non-space characters.
	Input string may contain leading or trailing spaces. However,
	your reversed string should not contain leading or trailing spaces.
	You need to reduce multiple spaces between two words to a single space
	in the reversed string.
*/

package lstring

func reverseWords2(s string) string {
	tmp := make([]byte, len(s)+1, len(s)+1)

	copyWord := func(start, end int, idx int) {
		for i := start; i < end; i++ {
			tmp[idx] = s[i]
			idx++
		}
		tmp[idx] = ' '
		idx++
	}

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}

		start, end := i, i+1
		for j := i; j >= 0 && s[j] != ' '; j-- {
			i--
		}
		start = i + 1

		copyWord(start, end, length)
		length += (end - start) + 1 // word + space
	}

	if length == 0 {
		return ""
	}
	tmp = tmp[:length-1]
	return string(tmp)
}
