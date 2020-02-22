package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximalRectangle(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, maximalRectangle([][]byte{{'0'}}))
	assert.Equal(1, maximalRectangle([][]byte{{'1'}}))

	input := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	assert.Equal(6, maximalRectangle(input))

	input = [][]byte{
		{'0', '1', '1', '0', '1'},
		{'1', '1', '0', '1', '0'},
		{'0', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0'},
	}
	assert.Equal(9, maximalRectangle(input))
}
