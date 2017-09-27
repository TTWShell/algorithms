/* https://leetcode.com/problems/spiral-matrix/description/
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

For example,
Given the following matrix:

[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
You should return [1,2,3,6,9,8,7,4,5].
*/

package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	length := m * n
	res, count := make([]int, length, length), 0
	rowStart, rowEnd, colStart, colEnd := 0, m-1, 0, n-1

	for count < length {
		for i := colStart; count < length && i <= colEnd; i++ {
			res[count] = matrix[rowStart][i]
			count++
		}
		rowStart++
		for j := rowStart; count < length && j <= rowEnd; j++ {
			res[count] = matrix[j][colEnd]
			count++
		}
		colEnd--
		for i := colEnd; count < length && i >= colStart; i-- {
			res[count] = matrix[rowEnd][i]
			count++
		}
		rowEnd--
		for j := rowEnd; count < length && j >= rowStart; j-- {
			res[count] = matrix[j][colStart]
			count++
		}
		colStart++
	}
	return res
}
