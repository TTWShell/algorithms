package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dominantIndex(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, dominantIndex([]int{3, 6, 1, 0}))
	assert.Equal(-1, dominantIndex([]int{1, 2, 3, 4}))
}
