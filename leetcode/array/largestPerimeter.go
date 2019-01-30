/* https://leetcode.com/problems/largest-perimeter-triangle/
Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.

Example 1:

	Input: [2,1,2]
	Output: 5

Example 2:

	Input: [1,2,1]
	Output: 0

Example 3:

	Input: [3,2,3,4]
	Output: 10

Example 4:

	put: [3,6,2,3]
	Output: 8

Note:

	3 <= A.length <= 10000
	1 <= A[i] <= 10^6
*/

package larray

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)

	for i := len(A) - 1; i > 1; i-- {
		a, b, c := A[i-2], A[i-1], A[i]
		if a+b <= c {
			continue
		}
		return a + b + c
	}
	return 0
}
