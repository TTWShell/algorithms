package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diStringMatch(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0, 4, 1, 3, 2}, diStringMatch("IDID"))
	assert.Equal([]int{0, 1, 2, 3}, diStringMatch("III"))
	assert.Equal([]int{3, 2, 0, 1}, diStringMatch("DDI"))
	assert.Equal([]int{0, 1, 3, 2}, diStringMatch("IID"))
}
