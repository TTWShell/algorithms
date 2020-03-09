package lbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, findMin2([]int{1, 3, 1, 1}))
	assert.Equal(1, findMin2([]int{1, 1, 3, 1}))
	assert.Equal(1, findMin2([]int{3, 3, 1, 3}))
	assert.Equal(1, findMin2([]int{3, 1, 3, 3, 3}))
	assert.Equal(1, findMin2([]int{1, 3, 5}))
	assert.Equal(0, findMin2([]int{2, 2, 2, 0, 1}))
	assert.Equal(1, findMin2([]int{3, 1, 3}))
	assert.Equal(0, findMin2([]int{0, 1, 2, 4, 5, 6, 7}))
	assert.Equal(0, findMin2([]int{4, 5, 6, 7, 0, 1, 2}))
}
