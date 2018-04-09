package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	assert := assert.New(t)

	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	assert.Equal([]int{3, 60}, numberOfLines(widths, "abcdefghijklmnopqrstuvwxyz"))

	widths = []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	assert.Equal([]int{2, 4}, numberOfLines(widths, "bbbcccdddaaa"))
}
