/* https://leetcode.com/problems/set-matrix-zeroes/description/
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

Follow up:
Did you use extra space?
A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

package larray

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	setColZero := func(matrix [][]int, col int) {
		for i := range matrix {
			matrix[i][col] = 0
		}
	}

	setRowZero := func(matrix [][]int, row, col int) {
		matrix[row][col] = 0 // cur
		// ago
		for j := col - 1; j >= 0; j-- {
			if matrix[row][j] == 0 &&
				((row-1 >= 0 && matrix[row-1][j] != 0) ||
					(row+1 < len(matrix) && matrix[row+1][j] != 0)) {
				break
			}
			matrix[row][j] = 0
		}
	}

	m, n := len(matrix), len(matrix[0])
	cols := make([]bool, 2, 2) // index --> col now-2, now-1 isNeedSetZero

	// init col index 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			cols[1] = true
			break
		}
	}
	// deal col index j
	for j := 1; j < n; j++ {
		// first deal col - 2
		if cols[0] {
			setColZero(matrix, j-2)
			cols[0] = false
		}
		cols[0], cols[1] = cols[1], cols[0]

		for i := 0; i < m; i++ {
			if matrix[i][j] == 0 && !cols[1] {
				cols[1] = true
			}
			// now can safe setZero row before curRow
			if matrix[i][j] == 0 || matrix[i][j-1] == 0 {
				setRowZero(matrix, i, j)
			}
		}
	}
	for i, flag := range cols {
		if col := n - 2 + i; flag && col >= 0 {
			setColZero(matrix, col)
		}
	}
}

/* O(m + n) space
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
*/
