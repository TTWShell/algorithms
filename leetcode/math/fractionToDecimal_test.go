package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fractionToDecimal(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("0.5", fractionToDecimal(1, 2))
	assert.Equal("2", fractionToDecimal(2, 1))
	assert.Equal("0.(6)", fractionToDecimal(2, 3))
	assert.Equal("-6.25", fractionToDecimal(-50, 8))
	assert.Equal("-0.58(3)", fractionToDecimal(7, -12))
}
