package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{
		{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"},
	})
}
