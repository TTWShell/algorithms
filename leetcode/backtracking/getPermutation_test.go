package lbacktracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getPermutation(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(getPermutation(1, 1), "1")
	assert.Equal(getPermutation(3, 1), "123")
	assert.Equal(getPermutation(3, 2), "132")
	assert.Equal(getPermutation(3, 3), "213")
	assert.Equal(getPermutation(3, 4), "231")
	assert.Equal(getPermutation(3, 5), "312")
	assert.Equal(getPermutation(3, 6), "321")
}
