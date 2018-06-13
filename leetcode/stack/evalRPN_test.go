package lstack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(9, evalRPN([]string{"2", "1", "+", "3", "*"}))
	assert.Equal(6, evalRPN([]string{"4", "13", "5", "/", "+"}))
	assert.Equal(3, evalRPN([]string{"4", "13", "5", "/", "+", "3", "-"}))
}
