/* https://leetcode.com/problems/sort-array-by-parity/
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.

Example 1:

	Input: [3,1,2,4]
	Output: [2,4,3,1]
	The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Note:

	1 <= A.length <= 5000
	0 <= A[i] <= 5000
*/

package larray

func sortArrayByParity(A []int) []int {
	res := make([]int, len(A), len(A))
	for i, start, end := 0, 0, len(A)-1; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[start] = A[i]
			start++
		} else {
			res[end] = A[i]
			end--
		}
	}
	return res
}
