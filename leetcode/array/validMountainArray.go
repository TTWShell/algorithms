/* https://leetcode.com/problems/valid-mountain-array/
Given an array A of integers, return true if and only if it is a valid mountain array.

Recall that A is a mountain array if and only if:

	A.length >= 3
	There exists some i with 0 < i < A.length - 1 such that:
	A[0] < A[1] < ... A[i-1] < A[i]
	A[i] > A[i+1] > ... > A[B.length - 1]

Example 1:

	Input: [2,1]
	Output: false

Example 2:

	Input: [3,5,5]
	Output: false

Example 3:

	Input: [0,3,2,1]
	Output: true

Note:

	0 <= A.length <= 10000
	0 <= A[i] <= 10000
*/

package larray

func validMountainArray(A []int) bool {
	if len(A) < 3 || A[0] >= A[1] {
		return false
	}

	top := false
	for i := 1; i < len(A); i++ {
		if (top == false && A[i-1] < A[i]) || (top == true && A[i-1] > A[i]) {
			continue
		}
		if A[i-1] == A[i] || top == true {
			return false
		}
		top = true
	}
	return top
}
