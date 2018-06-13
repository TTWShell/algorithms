package lunionfind

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solve(t *testing.T) {
	assert := assert.New(t)

	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	excepted := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	assert.Equal(excepted, board)

	board = [][]byte{}
	solve(board)
	assert.Equal([][]byte{}, board)

	board = [][]byte{{'O', 'O'}, {'O', 'O'}}
	solve(board)
	assert.Equal([][]byte{{'O', 'O'}, {'O', 'O'}}, board)
}
