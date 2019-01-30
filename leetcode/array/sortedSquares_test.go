package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1}, sortedSquares([]int{-1}))
	assert.Equal([]int{1}, sortedSquares([]int{1}))
	assert.Equal([]int{0, 4}, sortedSquares([]int{-2, 0}))
	assert.Equal([]int{0, 4}, sortedSquares([]int{0, 2}))
	assert.Equal([]int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
	assert.Equal([]int{4, 9, 9, 49, 121}, sortedSquares([]int{-7, -3, 2, 3, 11}))
}
