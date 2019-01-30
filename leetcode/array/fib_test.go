package larray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fib(0))
	assert.Equal(1, fib(1))
	assert.Equal(1, fib(2))
	assert.Equal(2, fib(3))
	assert.Equal(3, fib(4))
}
