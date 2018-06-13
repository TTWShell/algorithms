/* https://leetcode.com/problems/reverse-string-ii/#/description
Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string.
If there are less than k characters left, reverse all of them.
If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.

Example:
    Input: s = "abcdefg", k = 2
    Output: "bacdfeg"

Restrictions:
    The string consists of lower English letters only.
    Length of the given string and k will in the range [1, 10000]
*/

package lstring

func reverseStr(s string, k int) string {
	reverse := func(sList []byte, start, end int) {
		for i := 0; i < (end-start)/2; i++ {
			sList[start+i], sList[end-1-i] = sList[end-1-i], sList[start+i]
		}
	}

	sList := []byte(s)
	for i := 0; i < len(sList); i += 2 * k {
		end := i + k
		if end > len(sList) {
			end = len(sList)
		}
		reverse(sList, i, end)
	}
	return string(sList)
}
