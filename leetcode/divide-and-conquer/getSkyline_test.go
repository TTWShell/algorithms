package ldc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSkyline(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{}, getSkyline([][]int{}))

	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	excepted := [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}
	assert.Equal(excepted, getSkyline(buildings))
}
