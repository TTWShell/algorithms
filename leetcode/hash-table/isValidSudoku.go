/* https://leetcode.com/problems/valid-sudoku/#/description
Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

http://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png

A partially filled sudoku which is valid.

Note:
A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.
*/

package lht

func isValidSudoku(board [][]byte) bool {
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
