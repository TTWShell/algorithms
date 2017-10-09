package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTotal(t *testing.T) {
	assert := assert.New(t)

	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	assert.Equal(11, minimumTotal(triangle))
}
