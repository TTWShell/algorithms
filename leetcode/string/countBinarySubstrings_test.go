package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countBinarySubstrings(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, countBinarySubstrings("00110011"))
	assert.Equal(4, countBinarySubstrings("10101"))
}
