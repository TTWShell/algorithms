package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, canCompleteCircuit([]int{1, 3}, []int{1, 1}))
	assert.Equal(1, canCompleteCircuit([]int{1, 3}, []int{2, 2}))
	assert.Equal(-1, canCompleteCircuit([]int{3, 1}, []int{3, 2}))
	assert.Equal(-1, canCompleteCircuit([]int{3, 1, 10}, []int{3, 2, 10}))
}
