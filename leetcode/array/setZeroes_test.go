package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{}
	setZeroes(matrix)
	assert.Equal([][]int{}, matrix)

	input := [][]int{
		{1, 0, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	output := [][]int{
		{0, 0, 0},
		{4, 0, 6},
		{7, 0, 9},
	}
	setZeroes(input)
	assert.Equal(output, input)
}
