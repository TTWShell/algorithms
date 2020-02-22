package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	assert := assert.New(t)

	assert.False(isInterleave("a", "b", "abc"))

	assert.True(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.False(isInterleave("aabcc", "dbbca", "aadbbbaccc"))

	assert.True(isInterleave("", "", ""))
	assert.True(isInterleave("a", "", "a"))
	assert.True(isInterleave("", "a", "a"))
	assert.True(isInterleave("a", "b", "ab"))
	assert.True(isInterleave("b", "a", "ab"))
}
