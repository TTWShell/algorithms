package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_384(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	obj.Shuffle()
	assert.Equal([]int{1, 2, 3}, obj.Reset())
}
