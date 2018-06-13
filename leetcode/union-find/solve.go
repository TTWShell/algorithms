/* https://leetcode.com/problems/surrounded-regions/description/
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

For example,
X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
*/

package lunionfind

import "github.com/TTWShell/algorithms/data-structure/union-find"

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	row, col := len(board), len(board[0])
	uf := unionfind.New()
	charO, char1 := byte('O'), byte('1')

	var helper func(i, j int)
	helper = func(i, j int) {
		uf.Union(charO, [2]int{i, j})
		board[i][j] = char1
		if tmp := i - 1; tmp >= 0 && board[tmp][j] == charO {
			helper(tmp, j)
		}
		if tmp := i + 1; tmp < row && board[tmp][j] == charO {
			helper(tmp, j)
		}
		if tmp := j - 1; tmp >= 0 && board[i][tmp] == charO {
			helper(i, tmp)
		}
		if tmp := j + 1; tmp < col && board[i][tmp] == charO {
			helper(i, tmp)
		}
	}

	// check around only.
	lastRow, lastCol := row-1, col-1
	for j := 0; j < col; j++ {
		if board[0][j] == charO {
			helper(0, j)
		}
		if board[lastRow][j] == charO {
			helper(lastRow, j)
		}
	}
	for i := 1; i < row-1; i++ {
		if board[i][0] == charO {
			helper(i, 0)
		}
		if board[i][lastCol] == charO {
			helper(i, lastCol)
		}
	}

	for i := range board {
		for j := range board[i] {
			if !uf.Connected(charO, [2]int{i, j}) {
				board[i][j] = 'X'
			}
			if board[i][j] == char1 {
				board[i][j] = charO
			}
		}
	}
}
