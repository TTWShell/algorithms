package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subsets(t *testing.T) {
	assert := assert.New(t)

	expected := [][]int{
		{},
		{1}, {2}, {3},
		{1, 2}, {1, 3}, {2, 3},
		{1, 2, 3},
	}
	assert.Equal(expected, subsets([]int{1, 2, 3}))
}
