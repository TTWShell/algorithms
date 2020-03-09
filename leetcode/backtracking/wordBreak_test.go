package lbacktracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	assert := assert.New(t)

	excepted := []string{
		"cats and dog",
		"cat sand dog",
	}
	result := wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	assert.Subset(excepted, result)
	assert.Len(result, 2)

	excepted = []string{
		"pine apple pen apple",
		"pineapple pen apple",
		"pine applepen apple",
	}
	result = wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	assert.Subset(excepted, result)
	assert.Len(result, 3)

	assert.Equal([]string{}, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	result = wordBreak(s, []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})
	assert.Equal([]string{}, result)
}
