package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_largestPalindrome(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(9, largestPalindrome(1))
	assert.Equal(987, largestPalindrome(2))
	assert.Equal(123, largestPalindrome(3))
	assert.Equal(597, largestPalindrome(4))
	assert.Equal(677, largestPalindrome(5))
	assert.Equal(1218, largestPalindrome(6))
	assert.Equal(877, largestPalindrome(7))
	assert.Equal(475, largestPalindrome(8))
}
