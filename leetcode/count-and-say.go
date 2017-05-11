/* https://leetcode.com/problems/count-and-say/#/description
The count-and-say sequence is the sequence of integers beginning as follows:
1, 11, 21, 1211, 111221, ...

1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth sequence.

Note: The sequence of integers will be represented as a string.
*/

package leetcode

import "strconv"

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		r, ago, t := "", s[0], 1
		for j := 1; j < len(s); j++ {
			if ago == s[j] {
				t++
			} else {
				r += strconv.Itoa(t) + string(ago)
				t = 1
				ago = s[j]
			}
		}
		s = r + strconv.Itoa(t) + string(ago)
	}
	return s
}
