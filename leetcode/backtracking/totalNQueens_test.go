package lbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalNQueens(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, totalNQueens(0))
	assert.Equal(1, totalNQueens(1))
	assert.Equal(0, totalNQueens(2))
	assert.Equal(0, totalNQueens(3))
	assert.Equal(2, totalNQueens(4))
	assert.Equal(10, totalNQueens(5))
}
