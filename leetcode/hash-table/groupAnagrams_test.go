package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	assert := assert.New(t)

	output := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	assert.Equal(len(output), 3)
	for _, sub := range output {
		switch len(sub) {
		case 1:
			assert.Equal(sub, []string{"bat"})
		case 2:
			assert.Equal(sub, []string{"tan", "nat"})
		case 3:
			assert.Equal(sub, []string{"eat", "tea", "ate"})
		}
	}
}
