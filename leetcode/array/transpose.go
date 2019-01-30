/* https://leetcode.com/problems/transpose-matrix/
Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

Example 1:

	Input: [[1,2,3],[4,5,6],[7,8,9]]
	Output: [[1,4,7],[2,5,8],[3,6,9]]

Example 2:

	Input: [[1,2,3],[4,5,6]]
	Output: [[1,4],[2,5],[3,6]]

Note:

	1 <= A.length <= 1000
	1 <= A[0].length <= 1000
*/

package larray

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return [][]int{}
	}

	row, col := len(A), len(A[0])
	res := make([][]int, col, col)
	for i := 0; i < col; i++ {
		res[i] = make([]int, row, row)
		for j := 0; j < row; j++ {
			res[i][j] = A[j][i]
		}
	}
	return res
}
