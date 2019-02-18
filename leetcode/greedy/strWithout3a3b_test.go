package lgreedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strWithout3a3b(t *testing.T) {
	assert := assert.New(t)

	assert.Subset([]string{"ab", "ba"}, []string{strWithout3a3b(1, 1)})
	assert.Subset([]string{"abb", "bab", "bba"}, []string{strWithout3a3b(1, 2)})
	assert.Subset([]string{"baa", "aba", "aab"}, []string{strWithout3a3b(2, 1)})
	assert.Subset([]string{"bbab", "babb"}, []string{strWithout3a3b(1, 3)})
	assert.Subset([]string{"bbaa", "aabb", "abab", "baba"}, []string{strWithout3a3b(2, 2)})
	assert.Equal("bbabbab", strWithout3a3b(2, 5))
	assert.Equal("aabaa", strWithout3a3b(4, 1))
}
