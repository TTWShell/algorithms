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

type Board [][]byte

func (board Board) isValid(row, col int, c byte) bool {
	i, j := row/3*3, col/3*3
	for k := 0; k < 9; k++ {
		if board[row][k] == c {
			return false
		}
		if board[k][col] == c {
			return false
		}
		if board[i+k%3][j+k/3] == c {
			return false
		}
	}
	return true
}

func (board Board) Solve() bool {
	for row := range board {
		for col := range board[row] {
			if board[row][col] != '.' {
				continue
			}
			for _, num := range []byte("123456789") {
				if !board.isValid(row, col, num) {
					continue
				}
				board[row][col] = num
				if board.Solve() {
					return true
				} else {
					board[row][col] = '.'
				}
			}
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	newBoard := Board(board)
	newBoard.Solve()
}
