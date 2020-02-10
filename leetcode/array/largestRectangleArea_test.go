package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
