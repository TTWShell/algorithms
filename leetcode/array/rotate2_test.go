package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate2(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate2(matrix)
	assert.Equal(matrix, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate2(matrix)
	assert.Equal(matrix, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}})
}
