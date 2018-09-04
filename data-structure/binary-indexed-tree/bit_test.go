package bit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := New(nums)
	assert.Equal(b.len, len(nums)+1)
	assert.Equal(b.c, []int{0, 1, 3, 3, 10, 5, 11, 7, 36})
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := New(nums)
	b.Update(0, 9)
	assert.Equal(b.c, []int{0, 10, 12, 3, 19, 5, 11, 7, 45})
}

func TestSum(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := New(nums)
	assert.Equal(b.Sum(1), 1)
	assert.Equal(b.Sum(2), 3)
	assert.Equal(b.Sum(3), 6)
	assert.Equal(b.Sum(4), 10)
	assert.Equal(b.Sum(5), 15)
	assert.Equal(b.Sum(6), 21)
	assert.Equal(b.Sum(7), 28)
	assert.Equal(b.Sum(8), 36)
}
