package lmath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numSquares(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, numSquares(12))
	assert.Equal(2, numSquares(13))
}
