package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(findLengthOfLCIS([]int{1, 3, 5, 4, 7}), 3)
	assert.Equal(findLengthOfLCIS([]int{2, 2, 2, 2, 2}), 1)
	assert.Equal(findLengthOfLCIS([]int{1, 3, 5, 7}), 4)
}
