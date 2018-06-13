package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_exist(t *testing.T) {
	assert := assert.New(t)

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	assert.True(exist(board, "ABCCED"))
	assert.True(exist(board, "SEE"))
	assert.False(exist(board, "ABCB"))

	board = [][]byte{[]byte("ab"), []byte("cd")}
	assert.False(exist(board, "abcd"))
	assert.True(exist(board, "abdc"))
	assert.True(exist(board, "acdb"))
}
