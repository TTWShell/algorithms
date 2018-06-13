package lbs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, findMin([]int{1}))
	assert.Equal(0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	assert.Equal(4, findMin([]int{4, 5, 6, 7}))
	assert.Equal(2, findMin([]int{4, 5, 6, 7, 2}))
	assert.Equal(2, findMin([]int{4, 5, 6, 7, 8, 2, 3}))
}
