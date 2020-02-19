/* https://leetcode.com/problems/minimum-window-substring/
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

	Input: S = "ADOBECODEBANC", T = "ABC"
	Output: "BANC"

Note:

	If there is no such window in S that covers all characters in T, return the empty string "".
	If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

package lht

import (
	"math"
)

func minWindow(s string, t string) string {
	countT := make([]int, 128, 128)
	for i := range t {
		countT[t[i]]++
	}

	res := ""
	left, subCount, minLength := 0, 0, math.MaxInt32
	for i := 0; i < len(s); i++ {
		if countT[s[i]] >= 1 {
			subCount++
		}
		countT[s[i]]--

		for subCount == len(t) {
			if tmpLength := i - left + 1; minLength > tmpLength {
				minLength = tmpLength
				res = string(s[left : left+minLength])
			}
			if countT[s[left]] >= 0 {
				subCount--
			}
			countT[s[left]]++
			left++
		}
	}
	return res
}
