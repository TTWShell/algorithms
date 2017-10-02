package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	assert := assert.New(t)

	expected := [][]int{
		{},
		{1}, {2}, {3},
		{1, 2}, {1, 3}, {2, 3},
		{1, 2, 3},
	}
	assert.Equal(expected, subsetsWithDup([]int{1, 2, 3}))

	expected = [][]int{
		{},
		{1}, {2},
		{1, 2}, {2, 2},
		{1, 2, 2},
	}
	assert.Equal(expected, subsetsWithDup([]int{1, 2, 2}))

	expected = [][]int{
		{},
		{1}, {2},
		{1, 1}, {1, 2}, {2, 2},
		{1, 1, 2},
		{1, 2, 2},
		{1, 1, 2, 2},
	}
	assert.Equal(expected, subsetsWithDup([]int{1, 1, 2, 2}))

	expected = [][]int{
		{},
		{1}, {4},
		{1, 4}, {4, 4},
		{1, 4, 4}, {4, 4, 4},
		{1, 4, 4, 4}, {4, 4, 4, 4},
		{1, 4, 4, 4, 4},
	}
	assert.Equal(expected, subsetsWithDup([]int{4, 4, 4, 1, 4}))
}
