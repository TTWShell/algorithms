package lbs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	assert := assert.New(t)

	assert.False(searchMatrix([][]int{{1}, {3}}, 0))

	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	assert.True(searchMatrix(matrix, 3))
	assert.True(searchMatrix(matrix, 11))
	assert.True(searchMatrix(matrix, 34))

	assert.False(searchMatrix(matrix, 12))

	matrix = [][]int{
		{-100, -90, -85, -80},
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	assert.True(searchMatrix(matrix, -90))
	assert.True(searchMatrix(matrix, 3))
	assert.True(searchMatrix(matrix, 11))
	assert.True(searchMatrix(matrix, 50))
}
