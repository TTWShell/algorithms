package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKFrequent(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 1, 1, 2, 2, 3}
	assert.Equal([]int{1, 2}, topKFrequent(nums, 2))

	assert.Equal([]int{1}, topKFrequent([]int{1}, 1))
}
