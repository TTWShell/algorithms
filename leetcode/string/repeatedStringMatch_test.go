package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repeatedStringMatch(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, repeatedStringMatch("abcd", "cdabcdab"))
	assert.Equal(1, repeatedStringMatch("aa", "a"))
}
