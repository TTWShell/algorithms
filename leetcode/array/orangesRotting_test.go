package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_orangesRotting(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	assert.Equal(-1, orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	assert.Equal(0, orangesRotting([][]int{{0, 2}}))
	assert.Equal(-1, orangesRotting([][]int{{1, 1, 1, 1}}))
	assert.Equal(0, orangesRotting([][]int{{0}}))
	assert.Equal(-1, orangesRotting([][]int{{2}, {2}, {1}, {0}, {1}, {1}}))
}
