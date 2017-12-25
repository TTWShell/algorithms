package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compareVersion(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, compareVersion("13.37", "1.2"))
	assert.Equal(-1, compareVersion("1.2", "13.37"))
	assert.Equal(0, compareVersion("1.2", "1.2"))
	assert.Equal(-1, compareVersion("1", "1.1"))
	assert.Equal(0, compareVersion("1", "1.0"))
}
