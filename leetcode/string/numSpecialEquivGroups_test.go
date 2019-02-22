package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSpecialEquivGroups(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, numSpecialEquivGroups([]string{"a", "b", "c", "a", "c", "c"}))
	assert.Equal(4, numSpecialEquivGroups([]string{"aa", "bb", "ab", "ba"}))
	assert.Equal(3, numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
	assert.Equal(1, numSpecialEquivGroups([]string{"abcd", "cdab", "adcb", "cbad"}))
}
