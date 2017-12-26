package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestNumber(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("9534330", largestNumber([]int{3, 30, 34, 5, 9}))
}
