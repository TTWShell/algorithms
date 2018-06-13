/* https://leetcode.com/problems/rotate-string/description/
We are given two strings, A and B.

A shift on A consists of taking string A and moving the leftmost character to the rightmost position. For example, if A = 'abcde', then it will be 'bcdea' after one shift on A. Return True if and only if A can become B after some number of shifts on A.

Example 1:

    Input: A = 'abcde', B = 'cdeab'
    Output: true

Example 2:

    Input: A = 'abcde', B = 'abced'
    Output: false

Note:

    A and B will have length at most 100.
*/

package lstring

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	for i := 0; i < len(A); i++ {
		k, j := i, 0
		for k < len(A) && j < len(B) && A[k] == B[j] {
			k++
			j++
		}
		if k == len(A) && string(A[:i]) == string(B[j:]) {
			return true
		}
	}
	return false
}
