package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	assert := assert.New(t)

	assert.False(wordBreak("leetcode", []string{"code"}))
	assert.True(wordBreak("leetcode", []string{"leet", "code"}))
}
