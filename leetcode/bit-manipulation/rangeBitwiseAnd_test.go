package lbm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, rangeBitwiseAnd(5, 7))
	assert.Equal(0, rangeBitwiseAnd(0, 0))
}
