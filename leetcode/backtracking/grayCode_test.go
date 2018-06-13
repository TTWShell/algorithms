package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_grayCode(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]int{0}, grayCode(0))
	assert.Equal([]int{0, 1}, grayCode(1))
	assert.Equal([]int{0, 1, 3, 2}, grayCode(2))
	assert.Equal([]int{0, 1, 3, 2, 6, 7, 5, 4}, grayCode(3))
}
