package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotateString(t *testing.T) {
	assert := assert.New(t)

	assert.False(rotateString("abcde", "cdeab111"))
	assert.True(rotateString("abcde", "bcdea"))
	assert.True(rotateString("abcde", "cdeab"))
	assert.False(rotateString("abcde", "abced"))
	assert.True(rotateString("ckahkzpikz", "hkzpikzcka"))
}
