package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isOneBitCharacter(t *testing.T) {
	assert := assert.New(t)

	assert.True(isOneBitCharacter([]int{1, 0, 0}))
	assert.False(isOneBitCharacter([]int{1, 1, 1, 0}))
}
