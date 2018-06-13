package lstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress(t *testing.T) {
	assert := assert.New(t)

	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	assert.Equal(6, compress(chars))
	assert.Equal([]byte("a2b2c3"), chars[:6])

	chars = []byte{'a'}
	assert.Equal(1, compress(chars))
	assert.Equal([]byte("a"), chars[:1])

	chars = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	assert.Equal(4, compress(chars))
	assert.Equal([]byte("ab12"), chars[:4])

	chars = []byte{'G', 't', 'W', 'Y', 'v', '&', ':', 'd', '#', 'k'}
	assert.Equal(10, compress(chars))
	assert.Equal([]byte{'G', 't', 'W', 'Y', 'v', '&', ':', 'd', '#', 'k'}, chars[:10])
}
