package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_increasingTriplet(t *testing.T) {
	assert := assert.New(t)

	assert.True(increasingTriplet([]int{1, 2, 3, 4, 5}))
	assert.False(increasingTriplet([]int{5, 4, 3, 2, 1}))
	assert.True(increasingTriplet([]int{1, 3, 2, 4, 5}))
	assert.True(increasingTriplet([]int{2, 5, 3, 4, 5}))
	assert.True(increasingTriplet([]int{2, 5, 3, 4}))
	assert.False(increasingTriplet([]int{2, 1, 5, 0, 3}))
	assert.True(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	assert.True(increasingTriplet([]int{1, 2, -10, -8, -7}))
	assert.True(increasingTriplet([]int{0, 4, 1, -1, 2}))
	assert.True(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}
