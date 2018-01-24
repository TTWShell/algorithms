package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isToeplitzMatrix(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	assert.True(isToeplitzMatrix(matrix))

	matrix = [][]int{
		{1, 2},
		{2, 2},
	}
	assert.False(isToeplitzMatrix(matrix))

	matrix = [][]int{
		{36, 59, 71, 15, 26, 82, 87},
		{56, 36, 59, 71, 15, 26, 82},
		{15, 0, 36, 59, 71, 15, 26},
	}
	assert.False(isToeplitzMatrix(matrix))
}
