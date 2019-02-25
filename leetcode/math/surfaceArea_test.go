package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_surfaceArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(10, surfaceArea([][]int{{2}}))
	assert.Equal(34, surfaceArea([][]int{{1, 2}, {3, 4}}))
	assert.Equal(16, surfaceArea([][]int{{1, 0}, {0, 2}}))
	assert.Equal(32, surfaceArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	assert.Equal(46, surfaceArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}))
}
