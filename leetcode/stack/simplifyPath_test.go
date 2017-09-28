package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("/home", simplifyPath("/home/"))
	assert.Equal("/c", simplifyPath("/a/./b/../../c/"))
	assert.Equal("/", simplifyPath("/../"))
	assert.Equal("/home/foo", simplifyPath("/home//foo/"))
}
