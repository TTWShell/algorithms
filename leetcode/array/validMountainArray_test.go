package larray

import "testing"
import "github.com/stretchr/testify/assert"

func Test_validMountainArray(t *testing.T) {
	assert := assert.New(t)

	assert.False(validMountainArray([]int{2, 1}))
	assert.False(validMountainArray([]int{3, 5, 5}))
	assert.False(validMountainArray([]int{3, 4, 5}))
	assert.False(validMountainArray([]int{5, 4, 3}))
	assert.True(validMountainArray([]int{0, 3, 2, 1}))
}
