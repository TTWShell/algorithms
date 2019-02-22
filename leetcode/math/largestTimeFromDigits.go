/* https://leetcode.com/problems/largest-time-for-given-digits/description/
Given an array of 4 digits, return the largest 24 hour time that can be made.

The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from 00:00, a time is larger if more time has elapsed since midnight.

Return the answer as a string of length 5.  If no valid time can be made, return an empty string.

Example 1:

	Input: [1,2,3,4]
	Output: "23:41"

Example 2:

	Input: [5,5,5,5]
	Output: ""

Note:

	A.length == 4
	0 <= A[i] <= 9
*/

package lmath

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(A []int) string {
	sort.Ints(A)
	for i := 3; i >= 0; i-- {
		if A[i] > 2 {
			continue
		}
		for j := 3; j >= 0; j-- {
			if j == i || (A[i] == 2 && A[j] > 3) {
				continue
			}
			for m := 3; m >= 0; m-- {
				if m == i || m == j || A[m] > 5 {
					continue
				}
				return fmt.Sprintf("%d%d:%d%d", A[i], A[j], A[m], A[6-i-j-m])
			}
		}
	}
	return ""
}
