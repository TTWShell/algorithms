package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findShortestSubArray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, findShortestSubArray([]int{}))
	assert.Equal(1, findShortestSubArray([]int{1}))
	assert.Equal(2, findShortestSubArray([]int{1, 2, 2, 3, 1}))
	assert.Equal(6, findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
