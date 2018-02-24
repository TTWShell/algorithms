package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_computeArea(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(45, computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
