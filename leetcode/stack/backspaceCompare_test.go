package lstack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	assert := assert.New(t)

	assert.True(backspaceCompare("ab#c", "ad#c"))
	assert.True(backspaceCompare("ab##", "c#d#"))
	assert.True(backspaceCompare("a##c", "#a#c"))
	assert.False(backspaceCompare("a#c", "b"))
	assert.False(backspaceCompare("a#c", "ba"))
}
