package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"a1b2", "a1B2", "A1b2", "A1B2"}, letterCasePermutation("a1b2"))
	assert.Equal([]string{"3z4", "3Z4"}, letterCasePermutation("3z4"))
	assert.Equal([]string{"12345"}, letterCasePermutation("12345"))

	assert.Equal([]string{""}, letterCasePermutation(""))
	assert.Equal([]string{"3Z4", "3z4"}, letterCasePermutation("3Z4"))
}
