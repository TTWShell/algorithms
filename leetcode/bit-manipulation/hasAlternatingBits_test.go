package lbm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasAlternatingBits(t *testing.T) {
	assert := assert.New(t)

	assert.True(hasAlternatingBits(5))
	assert.False(hasAlternatingBits(7))
	assert.False(hasAlternatingBits(11))
	assert.True(hasAlternatingBits(10))
}
