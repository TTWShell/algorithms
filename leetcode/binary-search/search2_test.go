package lbs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search2(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{
		{3, 1, 1},
		{1, 1, 3},
		{1, 3, 1},

		{3, 1, 1, 1},
		{1, 3, 1, 1},
		{1, 1, 3, 1},
		{1, 1, 1, 3},

		{3, 1, 1, 1, 1},
		{1, 3, 1, 1, 1},
		{1, 1, 3, 1, 1},
		{1, 1, 1, 3, 1},
		{1, 1, 1, 1, 3},

		{4, 5, 6, 7, 7, 7, 0, 1, 1, 2},
	}
	for _, nums := range input {
		for _, num := range nums {
			assert.True(search2(nums, num), fmt.Sprintf("cur is %#v, %d", nums, num))
		}
	}

	nums := []int{4, 5, 6, 7, 7, 7, 0, 1, 1, 2}
	assert.False(search2(nums, 1000))
	assert.False(search2(nums, -1000))
}
