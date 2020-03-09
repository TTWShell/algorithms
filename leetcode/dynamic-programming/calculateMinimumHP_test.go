package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateMinimumHP(t *testing.T) {
	assert := assert.New(t)

	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	assert.Equal(7, calculateMinimumHP(dungeon))

	dungeon = [][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}
	assert.Equal(3, calculateMinimumHP(dungeon))
}
