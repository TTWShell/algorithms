package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortColors(t *testing.T) {
	assert := assert.New(t)

	nums := []int{2, 1, 0, 2, 1, 0, 1, 1, 0, 0, 2, 2}
	output := []int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}
	sortColors(nums)
	assert.Equal(output, nums)

	nums = []int{0, 0}
	sortColors(nums)
	assert.Equal([]int{0, 0}, nums)

	nums = []int{2, 1}
	sortColors(nums)
	assert.Equal([]int{1, 2}, nums)

	nums = []int{0, 2, 2}
	sortColors(nums)
	assert.Equal([]int{0, 2, 2}, nums)

	nums = []int{1, 2, 2, 2, 2, 0, 0, 0, 1, 1}
	sortColors(nums)
	assert.Equal([]int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2}, nums)
}
