package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transpose(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{}, transpose([][]int{}))
	assert.Equal([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}, transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	assert.Equal([][]int{{1, 4}, {2, 5}, {3, 6}}, transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
