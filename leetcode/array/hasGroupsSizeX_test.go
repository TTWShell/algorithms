package larray

import "testing"
import "github.com/stretchr/testify/assert"

func Test_hasGroupsSizeX(t *testing.T) {
	assert := assert.New(t)

	assert.True(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	assert.False(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	assert.False(hasGroupsSizeX([]int{1}))
	assert.True(hasGroupsSizeX([]int{1, 1}))
	assert.True(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
	assert.True(hasGroupsSizeX([]int{1, 1, 1, 1, 2, 2}))
	assert.False(hasGroupsSizeX([]int{0, 0, 0, 0, 0, 1, 1, 2, 3, 4}))
	assert.True(hasGroupsSizeX([]int{0, 0, 0, 1, 1, 1, 2, 2, 2}))
	assert.False(hasGroupsSizeX([]int{0, 0, 1}))
}
