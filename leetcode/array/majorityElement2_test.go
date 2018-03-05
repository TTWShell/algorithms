package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{}, majorityElement2([]int{}))
	assert.Equal([]int{2}, majorityElement2([]int{1, 1, 3, 4, 2, 2, 2}))
	assert.Equal([]int{2}, majorityElement2([]int{2, 2, 2, 1, 1, 3, 4, 5}))
	assert.Equal([]int{1}, majorityElement2([]int{1}))
}
