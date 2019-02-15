package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addToArrayForm(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2, 3, 4}, addToArrayForm([]int{1, 2, 0, 0}, 34))
	assert.Equal([]int{1, 2, 0, 0}, addToArrayForm([]int{1, 2, 0, 0}, 0))
	assert.Equal([]int{4, 5, 5}, addToArrayForm([]int{2, 7, 4}, 181))
	assert.Equal([]int{1, 0, 2, 1}, addToArrayForm([]int{2, 1, 5}, 806))
	assert.Equal([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1))
}
