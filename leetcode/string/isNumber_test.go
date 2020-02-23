package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	assert := assert.New(t)

	assert.True(isNumber("0"))
	assert.True(isNumber(" 0.1 "))
	assert.False(isNumber("abc"))
	assert.False(isNumber("1 a"))
	assert.True(isNumber("2e10"))
	assert.True(isNumber(" -90e3   "))
	assert.False(isNumber(" 1e"))
	assert.False(isNumber("e3"))
	assert.True(isNumber(" 6e-1"))
	assert.False(isNumber(" 99e2.5 "))
	assert.True(isNumber("53.5e93"))
	assert.False(isNumber(" --6 "))
	assert.False(isNumber("-+3"))
	assert.False(isNumber("95a54e53"))

	assert.False(isNumber("  "))
	assert.True(isNumber("3."))
	assert.False(isNumber("."))
	assert.False(isNumber(". 1"))
	assert.True(isNumber(".1"))
	assert.False(isNumber("4e+"))
	assert.False(isNumber(" -."))
}
