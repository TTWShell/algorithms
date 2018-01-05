package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	assert := assert.New(t)

	nums := []int{3, 2, 1, 5, 6, 4}
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	t.Log(sorted)
	for i := range sorted {
		assert.Equal(sorted[i], findKthLargest(nums, len(nums)-i), "k =", len(nums)-i)
	}

	nums = []int{3, 2, 1, -1, -2, -3}
	copy(sorted, nums)
	sort.Ints(sorted)

	t.Log(sorted)
	for i := range sorted {
		assert.Equal(sorted[i], findKthLargest(nums, len(nums)-i), "k =", len(nums)-i)
	}
}
