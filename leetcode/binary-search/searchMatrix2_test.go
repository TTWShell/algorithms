package lbs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchMatrix2(t *testing.T) {
	assert := assert.New(t)

	assert.False(searchMatrix2([][]int{}, 1))
	assert.True(searchMatrix2([][]int{{-5}}, -5))

	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	assert.True(searchMatrix2(matrix, 3))
	assert.False(searchMatrix2(matrix, -1))
	assert.False(searchMatrix2(matrix, 31))
	for i := range matrix {
		for j := range matrix[i] {
			assert.True(searchMatrix2(matrix, matrix[i][j]), fmt.Sprintf("target: %d", matrix[i][j]))
		}
	}
	assert.False(searchMatrix2(matrix, 20))

	matrix = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	for i := range matrix {
		for j := range matrix[i] {
			assert.True(searchMatrix2(matrix, matrix[i][j]), fmt.Sprintf("target: %d", matrix[i][j]))
		}
	}
}
