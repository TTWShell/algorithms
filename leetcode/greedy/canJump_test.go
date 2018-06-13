package lgreedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canJump(t *testing.T) {
	assert := assert.New(t)

	assert.True(canJump([]int{2, 3, 1, 1, 4}))
	assert.False(canJump([]int{3, 2, 1, 0, 4}))
	assert.False(canJump([]int{0, 2, 3}))
	assert.True(canJump([]int{0}))
}
