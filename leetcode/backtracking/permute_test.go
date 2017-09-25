package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(permute([]int{1, 2, 3}), [][]int{{1, 2, 3}, {1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1}})
}
