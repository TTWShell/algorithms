package lbit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NumArray(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 3, 5}
	obj := Constructor(nums)

	assert.Equal(9, obj.SumRange(0, 2))
	obj.Update(1, 2)
	assert.Equal(8, obj.SumRange(0, 2))
}
