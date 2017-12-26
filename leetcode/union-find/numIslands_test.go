package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numIslands(t *testing.T) {
	assert := assert.New(t)

	grid := [][]byte{
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	assert.Equal(1, numIslands(grid))

	grid = [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	assert.Equal(3, numIslands(grid))

	assert.Equal(1, numIslands([][]byte{{'1'}, {'1'}}))
}
