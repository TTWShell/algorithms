package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divisorGame(t *testing.T) {
	assert := assert.New(t)

	assert.False(divisorGame(1))
	assert.True(divisorGame(2))
	assert.False(divisorGame(3))
	assert.True(divisorGame(4))
}
