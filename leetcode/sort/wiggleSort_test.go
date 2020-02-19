package lsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wiggleSort(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	assert.Equal([]int{1, 6, 1, 5, 1, 4}, nums)

	nums = []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums)
	assert.Equal([]int{2, 3, 1, 3, 1, 2}, nums)
}
