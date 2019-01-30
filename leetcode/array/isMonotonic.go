/* https://leetcode.com/problems/monotonic-array/
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

Example 1:

	Input: [1,2,2,3]
	Output: true

Example 2:

	Input: [6,5,4,4]
	Output: true

Example 3:

	Input: [1,3,2]
	Output: false

Example 4:

	Input: [1,2,4,5]
	Output: true

Example 5:

	Input: [1,1,1]
	Output: true

Note:

	1 <= A.length <= 50000
	-100000 <= A[i] <= 100000

*/

package larray

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	i, flag := 0, false
	for i < len(A)-1 && A[i] == A[i+1] {
		i++
	}
	i++
	if i >= len(A) {
		return true // tail
	}
	if A[i-1] < A[i] {
		flag = true
	}

	for ; i < len(A); i++ {
		if A[i-1] == A[i] || (A[i-1] < A[i] && flag == true) || (A[i-1] > A[i] && flag == false) {
			continue
		}
		return false
	}
	return true
}
