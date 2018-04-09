package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_largestTriangleArea(t *testing.T) {
	assert := assert.New(t)

	assert.True(math.Abs(2.0-largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}})) < math.Pow10(-15))
}
