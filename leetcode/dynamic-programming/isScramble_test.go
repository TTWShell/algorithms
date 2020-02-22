package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isScramble(t *testing.T) {
	assert := assert.New(t)

	assert.True(isScramble("s", "s"))
	assert.False(isScramble("aa", "a"))
	assert.True(isScramble("sa", "as"))
	assert.False(isScramble("sa", "ab"))
	assert.True(isScramble("great", "rgeat"))
	assert.False(isScramble("abcde", "caebd"))
	assert.True(isScramble("abb", "bab"))
	assert.True(isScramble("abb", "bba"))

	assert.False(isScramble("aacaacccacbcbcbcbb", "bcacabbbaaabcccccc"))
}
