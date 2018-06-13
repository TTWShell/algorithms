package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isAdditiveNumber(t *testing.T) {
	assert := assert.New(t)

	assert.True(isAdditiveNumber("112358"))
	assert.True(isAdditiveNumber("199100199"))
	assert.False(isAdditiveNumber("1023"))
	assert.False(isAdditiveNumber("1203"))
	assert.True(isAdditiveNumber("000"))
	assert.False(isAdditiveNumber("199001200"))
}
