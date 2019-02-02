package lbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, peakIndexInMountainArray([]int{0, 1, 0}))
	assert.Equal(1, peakIndexInMountainArray([]int{0, 2, 1, 0}))
	assert.Equal(2, peakIndexInMountainArray([]int{0, 2, 3, 0}))
}
