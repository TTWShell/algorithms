/* https://leetcode.com/problems/game-of-life/description/
According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

    1. Any live cell with fewer than two live neighbors dies, as if caused by under-population.
	当前细胞为存活状态时，当周围低于2个（不包含2个）存活细胞时， 该细胞变成死亡状态。（模拟生命数量稀少）
    2. Any live cell with two or three live neighbors lives on to the next generation.
	当前细胞为存活状态时，当周围有2个或3个存活细胞时， 该细胞保持原样。
    3. Any live cell with more than three live neighbors dies, as if by over-population.
	当前细胞为存活状态时，当周围有3个以上的存活细胞时，该细胞变成死亡状态。（模生命数量过多）
    4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	当前细胞为死亡状态时，当周围有3个存活细胞时，该细胞变成存活状态。 （模拟繁殖）

Write a function to compute the next state (after one update) of the board given its current state.

Follow up:

    1. Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells first and then use their updated values to update other cells.
    2. In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array. How would you address these problems?
*/

package leetcode

func gameOfLife(board [][]int) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])
	if col == 0 {
		return
	}

	const (
		DIE  = 0
		LIVE = 1
	)

	helper := func(curRow, curCol int) int {
		count := 0 // live neighbors
		for i := curRow - 1; i <= curRow+1; i++ {
			for j := curCol - 1; j <= curCol+1; j++ {
				if 0 <= i && i < row && 0 <= j && j < col && !(i == curRow && j == curCol) && board[i][j] == LIVE {
					count++
				}
			}
		}
		if count == 3 || (board[curRow][curCol] == LIVE && count == 2) {
			return LIVE
		}
		return DIE
	}

	tmp1, tmp2 := make([]int, row), make([]int, row)
	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			tmp2[i] = helper(i, j)
		}
		if j > 0 {
			for i := 0; i < row; i++ {
				board[i][j-1] = tmp1[i]
			}
		}
		tmp1, tmp2 = tmp2, tmp1
	}
	for i := 0; i < row; i++ {
		board[i][col-1] = tmp1[i]
	}
}
