package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([][]int{{1, 2, 4}}, combinationSum3(3, 7))
	assert.Equal([][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}, combinationSum3(3, 9))
}
