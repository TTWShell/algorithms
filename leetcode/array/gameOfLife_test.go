package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDiffInBST(t *testing.T) {
	assert := assert.New(t)

	board := [][]int{}
	gameOfLife(board)
	assert.Equal([][]int{}, board)

	board = [][]int{{}}
	gameOfLife(board)
	assert.Equal([][]int{{}}, board)

	board = [][]int{{1}}
	gameOfLife(board)
	assert.Equal([][]int{{0}}, board)

	board = [][]int{{0, 1}}
	gameOfLife(board)
	assert.Equal([][]int{{0, 0}}, board)

	board = [][]int{{1, 0}}
	gameOfLife(board)
	assert.Equal([][]int{{0, 0}}, board)

	board = [][]int{{1, 1}}
	gameOfLife(board)
	assert.Equal([][]int{{0, 0}}, board)

	board = [][]int{{1, 1}, {1, 0}}
	gameOfLife(board)
	assert.Equal([][]int{{1, 1}, {1, 1}}, board)
}
