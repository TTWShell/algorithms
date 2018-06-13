package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(permuteUnique([]int{1, 1, 2}), [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}})
}
