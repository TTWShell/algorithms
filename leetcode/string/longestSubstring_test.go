package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestSubstring(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, longestSubstring("aaabb", 3))
	assert.Equal(5, longestSubstring("ababbc", 2))
	assert.Equal(6, longestSubstring("aaabbb", 3))
	assert.Equal(0, longestSubstring("weitong", 2))
}
