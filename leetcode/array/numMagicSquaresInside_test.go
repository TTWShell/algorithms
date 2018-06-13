package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numMagicSquaresInside(t *testing.T) {
	assert := assert.New(t)

	input := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	assert.Equal(1, numMagicSquaresInside(input))
	input[2][3] = 12
	assert.Equal(1, numMagicSquaresInside(input))
}
