package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{}
	setZeroes(matrix)
	assert.Equal([][]int{}, matrix)

	matrix = [][]int{{0}}
	setZeroes(matrix)
	assert.Equal([][]int{{0}}, matrix)

	matrix = [][]int{{1, 0}}
	setZeroes(matrix)
	assert.Equal([][]int{{0, 0}}, matrix)

	matrix = [][]int{{0, 1}}
	setZeroes(matrix)
	assert.Equal([][]int{{0, 0}}, matrix)

	matrix = [][]int{{0}, {1}}
	setZeroes(matrix)
	assert.Equal([][]int{{0}, {0}}, matrix)

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

	input = [][]int{
		{0, 0, 0, 5},
		{4, 3, 1, 4},
		{0, 1, 1, 4},
		{1, 2, 1, 3},
		{0, 0, 1, 1},
	}
	output = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 4},
		{0, 0, 0, 0},
		{0, 0, 0, 3},
		{0, 0, 0, 0},
	}
	setZeroes(input)
	assert.Equal(output, input)

	input = [][]int{
		{1, 2, 3, 4},
		{5, 0, 5, 2},
		{8, 9, 2, 0},
		{5, 7, 2, 1},
	}
	output = [][]int{
		{1, 0, 3, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{5, 0, 2, 0},
	}
	setZeroes(input)
	assert.Equal(output, input)
}
