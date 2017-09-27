package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(generateMatrix(3), [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}})
}
