package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_projectionArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, projectionArea([][]int{{2}}))
	assert.Equal(17, projectionArea([][]int{{1, 2}, {3, 4}}))
	assert.Equal(8, projectionArea([][]int{{1, 0}, {0, 2}}))
	assert.Equal(14, projectionArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	assert.Equal(21, projectionArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}))
}
