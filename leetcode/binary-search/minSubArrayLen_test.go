package lbs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, minSubArrayLen(100, []int{}))
	assert.Equal(2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(5, minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}
