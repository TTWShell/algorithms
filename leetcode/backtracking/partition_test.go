package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]string{{"aa", "b"}, {"a", "a", "b"}}, partition("aab"))
	assert.Equal([][]string{{"bb"}, {"b", "b"}}, partition("bb"))
}
