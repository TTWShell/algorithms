package larray

import "testing"
import "github.com/stretchr/testify/assert"

func Test_fairCandySwap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{1, 2}, fairCandySwap([]int{1, 1}, []int{2, 2}))
	assert.Equal([]int{2, 1}, fairCandySwap([]int{2, 2}, []int{1, 1}))
}
