package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
