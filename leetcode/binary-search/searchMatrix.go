/* https://leetcode.com/problems/search-a-2d-matrix/description/
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
For example,

Consider the following matrix:

[
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
Given target = 3, return true.
*/

package lbs

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	switch m {
	case 0:
		return false
	case 1:
		start, end := 0, len(matrix[0])-1
		for start <= end {
			mid := start + (end-start)/2
			if matrix[0][mid] == target {
				return true
			} else if matrix[0][mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	default:
		n := len(matrix[0]) - 1
		start, end := 0, m-1
		for start <= end {
			mid := start + (end-start)/2
			if target < matrix[mid][0] {
				end = mid - 1
			} else if matrix[mid][n] < target {
				start = mid + 1
			} else {
				return searchMatrix([][]int{matrix[mid]}, target)
			}
		}
	}
	return false
}
