package lheap

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, findKthLargest([]int{1, 2, 3, 4, 5, 6}, 1))
	assert.Equal(99, findKthLargest([]int{99, 99}, 1))

	nums := []int{3, 2, 1, 5, 6, 4}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	for i := range sorted {
		assert.Equal(sorted[i], findKthLargest(nums, len(nums)-i), "k =", len(nums)-i)
	}

	nums = []int{3, 2, 1, -1, -2, -3}
	copy(sorted, nums)
	sort.Ints(sorted)

	for i := range sorted {
		assert.Equal(sorted[i], findKthLargest(nums, len(nums)-i), "k =", len(nums)-i)
	}
}
