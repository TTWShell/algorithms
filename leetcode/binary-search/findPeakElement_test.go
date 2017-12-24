package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, findPeakElement([]int{1, 2, 3, 1}))
}
