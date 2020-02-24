package lbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{}, solveNQueens(0))
	assert.Equal([][]string{{"Q"}}, solveNQueens(1))
	assert.Equal([][]string{}, solveNQueens(2))
	assert.Equal([][]string{}, solveNQueens(3))

	excepted := [][]string{
		{
			".Q..", // Solution 1
			"...Q",
			"Q...",
			"..Q.",
		},
		{
			"..Q.", // Solution 2
			"Q...",
			"...Q",
			".Q..",
		},
	}
	assert.Equal(excepted, solveNQueens(4))

	assert.Equal(10, len(solveNQueens(5)))
}
