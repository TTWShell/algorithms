/* https://leetcode.com/problems/reverse-string/#/description
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".
*/

package lstring

func reverseString(s string) string {
	r := []rune(s)
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}
	return string(r)
}
