/* https://leetcode.com/problems/sudoku-solver/#/description
Write a program to solve a Sudoku puzzle by filling the empty cells.

Empty cells are indicated by the character '.'.

You may assume that there will be only one unique solution.

http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png

A sudoku puzzle...

http://upload.wikimedia.org/wikipedia/commons/thumb/3/31/Sudoku-by-L2G-20050714_solution.svg/250px-Sudoku-by-L2G-20050714_solution.svg.png

..and its solution numbers marked in red.
*/

package leetcode

func solveSudoku(board [][]byte) {
	isValidSudoku := func(board [][]byte) bool {
		var rowMask, colMask, areaMask [9][9]bool
		for r := range board {
			for c := range board[r] {
				if board[r][c] != '.' {
					digit := board[r][c] - '0' - 1
					area := 3*(r/3) + c/3
					if rowMask[r][digit] || colMask[c][digit] || areaMask[area][digit] {
						return false
					}
					rowMask[r][digit], colMask[c][digit], areaMask[area][digit] = true, true, true
				}
			}
		}
		return true
	}

	var helper func(board [][]byte) bool
	helper = func(board [][]byte) bool {
		for row := range board {
			for col := range board[row] {
				if board[row][col] == '.' {
					for _, num := range []byte("123456789") {
						board[row][col] = num
						if !isValidSudoku(board) {
							board[row][col] = '.'
						} else {
							if helper(board) {
								return true
							} else {
								board[row][col] = '.'
							}
						}
					}
					return false
				}
			}
		}
		return true
	}

	helper(board)
}
