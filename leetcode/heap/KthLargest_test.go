package lheap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_KthLargest(t *testing.T) {
	assert := assert.New(t)

	obj := Constructor(3, []int{4, 5, 8, 2})

	assert.Equal(4, obj.Add(3))
	assert.Equal(5, obj.Add(5))
	assert.Equal(5, obj.Add(10))
	assert.Equal(8, obj.Add(9))
	assert.Equal(8, obj.Add(4))

	obj = Constructor(1, []int{})
	assert.Equal(-3, obj.Add(-3))
	assert.Equal(-2, obj.Add(-2))
	assert.Equal(-2, obj.Add(-4))
	assert.Equal(0, obj.Add(0))
	assert.Equal(4, obj.Add(4))

	obj = Constructor(3, []int{5, -1})
	assert.Equal(-1, obj.Add(2))
	assert.Equal(1, obj.Add(1))
	assert.Equal(1, obj.Add(-1))
	assert.Equal(2, obj.Add(3))
	assert.Equal(3, obj.Add(4))

	obj = Constructor(3, []int{1, 1})
	assert.Equal(1, obj.Add(1))
	assert.Equal(1, obj.Add(1))
	assert.Equal(1, obj.Add(3))
	assert.Equal(1, obj.Add(3))
	assert.Equal(3, obj.Add(3))
	assert.Equal(3, obj.Add(4))
	assert.Equal(3, obj.Add(4))
	assert.Equal(4, obj.Add(4))

	obj = Constructor(2, []int{0})
	assert.Equal(-1, obj.Add(-1))
	assert.Equal(0, obj.Add(1))
	assert.Equal(0, obj.Add(-2))
	assert.Equal(0, obj.Add(-4))
	assert.Equal(1, obj.Add(3))

	obj = Constructor(7, []int{-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1})
	assert.Equal(3, obj.Add(2))
}
