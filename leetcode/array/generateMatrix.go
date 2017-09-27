/* https://leetcode.com/problems/spiral-matrix-ii/description/
Given an integer n, generate a square matrix filled with elements from 1 to n^2 in spiral order.

For example,
Given n = 3,

You should return the following matrix:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

package leetcode

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n, n)
	}

	count := 1
	rowStart, rowEnd, colStart, colEnd := 0, n-1, 0, n-1
	for count <= n*n {
		for i := colStart; i <= colEnd; i++ {
			matrix[rowStart][i] = count
			count++
		}
		rowStart++
		for j := rowStart; j <= rowEnd; j++ {
			matrix[j][colEnd] = count
			count++
		}
		colEnd--
		for i := colEnd; i >= colStart; i-- {
			matrix[rowEnd][i] = count
			count++
		}
		rowEnd--
		for j := rowEnd; j >= rowStart; j-- {
			matrix[j][colStart] = count
			count++
		}
		colStart++
	}
	return matrix
}
