package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isRectangleOverlap(t *testing.T) {
	assert := assert.New(t)

	assert.True(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	assert.True(isRectangleOverlap([]int{1, 1, 3, 3}, []int{0, 0, 2, 2}))
	assert.False(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
	assert.False(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 3, 3, 3}))
}
