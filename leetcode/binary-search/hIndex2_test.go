package lbs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDiffInBST(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, hIndex([]int{}))
	assert.Equal(3, hIndex([]int{0, 1, 3, 5, 6}))
	assert.Equal(1, hIndex([]int{100}))
	assert.Equal(2, hIndex([]int{11, 12}))
	assert.Equal(2, hIndex([]int{0, 0, 4, 4}))
	assert.Equal(3, hIndex([]int{1, 4, 7, 9}))
}
