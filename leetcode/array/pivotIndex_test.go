package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(-1, pivotIndex([]int{1, 2, 3}))
}
