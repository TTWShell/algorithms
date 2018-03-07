package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber3(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{5, 3}, singleNumber3([]int{1, 2, 1, 3, 2, 5}))
}
