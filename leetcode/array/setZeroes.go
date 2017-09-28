/* https://leetcode.com/problems/set-matrix-zeroes/description/
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

Follow up:
Did you use extra space?
A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

package leetcode

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rows := make([]bool, m, m)
	cols := make([]bool, n, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if rows[i] {
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < n; j++ {
		if cols[j] {
			for i := range matrix {
				matrix[i][j] = 0
			}
		}
	}
}
