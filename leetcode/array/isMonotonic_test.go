package larray

import "testing"
import "github.com/stretchr/testify/assert"

func Test_isMonotonic(t *testing.T) {
	assert := assert.New(t)

	assert.True(isMonotonic([]int{1}))
	assert.True(isMonotonic([]int{1, 2, 2, 3}))
	assert.True(isMonotonic([]int{6, 5, 4, 4}))
	assert.False(isMonotonic([]int{1, 3, 2}))
	assert.True(isMonotonic([]int{1, 2, 4, 5}))
	assert.True(isMonotonic([]int{1, 1, 1}))
}
