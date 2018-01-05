package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rob2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, rob2([]int{}))
	assert.Equal(2, rob2([]int{2}))

	input := [][]int{[]int{1, 2, 3}, []int{1, 2, 5, 7, 4, 3, 1, 3, 11, 3, 4, 2, 1, 8, 4, 11, 2, 2, 34}}
	result := []int{3, 80}
	for i := 0; i < len(input); i++ {
		assert.Equal(result[i], rob2(input[i]))
	}
}
