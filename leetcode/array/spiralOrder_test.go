package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(spiralOrder([][]int{{2, 3}}), []int{2, 3})
	assert.Equal(spiralOrder([][]int{}), []int{})
	assert.Equal(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
}
