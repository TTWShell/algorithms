package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	assert := assert.New(t)
	nums := []int{4, 5, 6, 7, -10, -5, -4, 0, 1, 2}

	for i, target := range nums {
		assert.Equal(i, search(nums, target))
	}

	assert.Equal(-1, search(nums, 1000))
	assert.Equal(-1, search(nums, -1000))
}
