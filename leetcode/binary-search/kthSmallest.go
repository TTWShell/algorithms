/* https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 ≤ k ≤ n2.
*/

package lbs

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	lo, hi := matrix[0][0], matrix[n-1][n-1]+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		j := n - 1
		for i := 0; i < n; i++ {
			for j >= 0 && mid < matrix[i][j] {
				j--
			}
			count += j + 1
		}
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
