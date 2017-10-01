/* https://leetcode.com/problems/word-search/description/
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

For example,
Given board =

[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
word = "ABCCED", -> returns true,
word = "SEE", -> returns true,
word = "ABCB", -> returns false.
*/

package leetcode

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	words := []byte(word)
	m, n := len(board), len(board[0])

	var helper func(row, col int, words []byte) bool
	helper = func(row, col int, words []byte) bool {
		if board[row][col] != words[0] {
			return false
		}
		if len(words) == 1 {
			return true
		}

		board[row][col] = board[row][col] - 26
		neighbors := [][]int{{row - 1, col}, {row, col - 1}, {row, col + 1}, {row + 1, col}}
		for _, neighbor := range neighbors {
			i, j := neighbor[0], neighbor[1]
			if i < 0 || i >= m || j < 0 || j >= n ||
				board[i][j] != words[1] {
				continue
			}
			if helper(i, j, words[1:]) {
				board[row][col] = board[row][col] + 26
				return true
			}
		}
		board[row][col] = board[row][col] + 26
		return false
	}

	for i := range board {
		for j := range board[i] {
			if helper(i, j, words) {
				return true
			}
		}
	}
	return false
}
