package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates2(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(5, removeDuplicates2(nums))
	assert.Equal([]int{1, 1, 2, 2, 3, 3}, nums)

	nums = []int{1, 1, 1, 2, 2, 3, 3, 3}
	assert.Equal(6, removeDuplicates2(nums))
	assert.Equal([]int{1, 1, 2, 2, 3, 3, 3, 3}, nums)

	nums = []int{1}
	assert.Equal(1, removeDuplicates2(nums))
	assert.Equal([]int{1}, nums)

	nums = []int{1, 1, 1, 1}
	assert.Equal(2, removeDuplicates2(nums))
	assert.Equal([]int{1, 1, 1, 1}, nums)

	nums = []int{}
	assert.Equal(0, removeDuplicates2(nums))
	assert.Equal([]int{}, nums)
}
