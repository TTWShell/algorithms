package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestRangeI(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, smallestRangeI([]int{1}, 0))
	assert.Equal(6, smallestRangeI([]int{0, 10}, 2))
	assert.Equal(0, smallestRangeI([]int{1, 3, 6}, 3))
	assert.Equal(0, smallestRangeI([]int{9, 9, 2, 8, 7}, 4))
	assert.Equal(916, smallestRangeI([]int{3021, 654, 5072, 9812, 4636, 3970, 2381, 1979, 9794, 4032}, 4121))
}
