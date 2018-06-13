package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SumRegion(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	numMatrix := NumMatrixConstructor(matrix)
	assert.Equal(8, numMatrix.SumRegion(2, 1, 4, 3))
	assert.Equal(11, numMatrix.SumRegion(1, 1, 2, 2))
	assert.Equal(12, numMatrix.SumRegion(1, 2, 2, 4))

	assert.Equal(3, numMatrix.SumRegion(0, 0, 0, 0))
	assert.Equal(58, numMatrix.SumRegion(0, 0, 4, 4))
	assert.Equal(4, numMatrix.SumRegion(0, 0, 0, 2))
}
