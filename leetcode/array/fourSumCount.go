/* https://leetcode.com/problems/4sum-ii/
Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500.
All integers are in the range of -2^28 to 2^28 - 1 and the result is guaranteed to be at most 2^31 - 1.

Example:

	Input:
	A = [ 1, 2]
	B = [-2,-1]
	C = [-1, 2]
	D = [ 0, 2]

	Output:
	2

	Explanation:
	The two tuples are:
	1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
	2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

package larray

func fourSumCount(A []int, B []int, C []int, D []int) int {
	AB := make(map[int]int, len(A)*len(A))
	for i := range A {
		for j := range B {
			AB[A[i]+B[j]]++
		}
	}

	res := 0
	for i := range C {
		for j := range D {
			if v, ok := AB[-C[i]-D[j]]; ok {
				res += v
			}
		}
	}
	return res
}
