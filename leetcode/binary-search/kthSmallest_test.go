package lbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	assert := assert.New(t)

	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}

	assert.Equal(13, kthSmallest(matrix, 8))
}
