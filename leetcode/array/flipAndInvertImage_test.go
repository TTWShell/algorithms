package larray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	assert := assert.New(t)
	input := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	output := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	assert.Equal(output, flipAndInvertImage(input))

	input = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	output = [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}
	assert.Equal(output, flipAndInvertImage(input))

	assert.Equal([][]int{}, flipAndInvertImage([][]int{}))
}
