package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	assert := assert.New(t)

	assert.True(isSubsequence("", "ahbgdc"))
	assert.True(isSubsequence("abc", "ahbgdc"))
	assert.False(isSubsequence("axc", "ahbgdc"))

	assert.True(isSubsequence("ace", "abcde"))
	assert.False(isSubsequence("aec", "abcde"))
}
