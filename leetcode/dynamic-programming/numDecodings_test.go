package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numDecodings(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, numDecodings(""))
	assert.Equal(0, numDecodings("0"))
	assert.Equal(1, numDecodings("1"))
	assert.Equal(0, numDecodings("01"))
	assert.Equal(1, numDecodings("10"))
	assert.Equal(2, numDecodings("12"))
	assert.Equal(1, numDecodings("27"))
	assert.Equal(0, numDecodings("100"))
	assert.Equal(1, numDecodings("110"))
	assert.Equal(5, numDecodings("1212"))
}
