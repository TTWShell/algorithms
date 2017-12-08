package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, singleNumber2([]int{0, 0, 0, 3}))
	assert.Equal(3, singleNumber2([]int{2, 2, 2, 3}))
}
