package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("", reverseWords2(""))
	assert.Equal("blue is sky the", reverseWords2("the sky is blue"))
	assert.Equal("world! hello", reverseWords2("  hello world!  "))
	assert.Equal("example good a", reverseWords2("a good   example"))
}
