package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, minPathSum([][]int{{0}}))
	assert.Equal(1, minPathSum([][]int{{1}}))
	assert.Equal(3, minPathSum([][]int{{1, 2}}))
	assert.Equal(3, minPathSum([][]int{{1}, {2}}))
	assert.Equal(7, minPathSum([][]int{{1, 3, 2}, {2, 4, 1}}))
}
