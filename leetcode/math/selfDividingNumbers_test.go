package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(selfDividingNumbers(1, 22), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22})
}
