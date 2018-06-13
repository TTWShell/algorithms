package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	assert := assert.New(t)

	assert.False(validPalindrome("abc"))
	assert.True(validPalindrome("aba"))
	assert.True(validPalindrome("abca"))
	assert.False(validPalindrome("abcda"))
	assert.True(validPalindrome("ebcbbececabbacecbbcbe"))
}
