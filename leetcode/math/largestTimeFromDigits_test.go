package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestTimeFromDigits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("00:00", largestTimeFromDigits([]int{0, 0, 0, 0}))
	assert.Equal("10:00", largestTimeFromDigits([]int{0, 1, 0, 0}))
	assert.Equal("23:41", largestTimeFromDigits([]int{1, 2, 3, 4}))
	assert.Equal("", largestTimeFromDigits([]int{5, 5, 5, 5}))
}
