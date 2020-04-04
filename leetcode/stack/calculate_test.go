package lstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(2, calculate("1 + 1"))
	assert.Equal(3, calculate(" 2-1 + 2 "))
	assert.Equal(23, calculate("(1+(4+5+2)-3)+(6+8)"))
}
