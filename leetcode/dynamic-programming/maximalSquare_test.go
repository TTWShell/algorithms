package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	assert.Equal(4, maximalSquare(matrix))

	matrix = [][]byte{
		[]byte("101001110"),
		[]byte("111000001"),
		[]byte("001100011"),
		[]byte("011001001"),
		[]byte("110110010"),
		[]byte("011111101"),
		[]byte("101110010"),
		[]byte("111010001"),
		[]byte("011110010"),
		[]byte("100111000"),
	}
	assert.Equal(4, maximalSquare(matrix))

	matrix = [][]byte{}
	assert.Equal(0, maximalSquare(matrix))
}
