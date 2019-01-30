package lbm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(uint32(964176192), reverseBits(43261596))
	assert.Equal(uint32(3221225471), reverseBits(4294967293))
}
