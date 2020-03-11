package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestPalindrome(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abba", shortestPalindrome("abba"))
	assert.Equal("aaacecaaa", shortestPalindrome("aacecaaa"))
	assert.Equal("dcbabcd", shortestPalindrome("abcd"))
}
