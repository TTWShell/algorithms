package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculate2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(7, calculate2("2*2+3"))
	assert.Equal(1, calculate2("1-1+1"))
	assert.Equal(9, calculate2("2+3+4"))
	assert.Equal(1119, calculate2("12+103+1004"))

	assert.Equal(7, calculate2("3+2*2"))
	assert.Equal(1, calculate2(" 3/2 "))
	assert.Equal(5, calculate2(" 3+5 / 2 "))
}
