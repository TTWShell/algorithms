/* https://leetcode.com/problems/shortest-distance-to-a-character/description/
Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

Example 1:

    Input: S = "loveleetcode", C = 'e'
    Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]

Note:

    S string length is in [1, 10000].
    C is a single character, and guaranteed to be in string S.
    All letters in S and C are lowercase.
*/

package lstring

func shortestToChar(S string, C byte) []int {
	left, res := -1, make([]int, len(S))

	helper := func(left, cur int) {
		if left == -1 {
			// head
			for i := cur - 1; i >= 0; i-- {
				res[i] = res[i+1] + 1
			}
		} else if cur == len(S) {
			// tail
			for i := left + 1; i < len(S); i++ {
				res[i] = res[i-1] + 1
			}
		} else {
			// mid
			for l, r := left+1, cur-1; r >= l && l >= 0 && r < len(S); l, r = l+1, r-1 {
				res[l] = res[l-1] + 1
				res[r] = res[r+1] + 1
			}
		}
	}

	for i := 0; i < len(S); i++ {
		if S[i] == C {
			helper(left, i)
			left = i
		} else if i == len(S)-1 {
			helper(left, len(S))
		}
	}
	return res
}
