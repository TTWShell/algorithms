package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combine(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(combine(4, 0))
	assert.Empty(combine(0, 0))
	assert.Empty(combine(1, 2))

	assert.Equal([][]int{{1}, {2}, {3}, {4}}, combine(4, 1))
	assert.Equal([][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine(4, 2))
	assert.Equal([][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}, combine(4, 3))
	assert.Equal([][]int{[]int{1, 2, 3, 4}}, combine(4, 4))

}
