package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	assert.Equal(0, uniquePathsWithObstacles([][]int{{1}}))
	assert.Equal(1, uniquePathsWithObstacles([][]int{{0}}))
	assert.Equal(0, uniquePathsWithObstacles([][]int{{0, 1}}))
	assert.Equal(1, uniquePathsWithObstacles([][]int{{0}, {0}}))
}
