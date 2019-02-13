package lbm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hammingWeight(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, hammingWeight(11))
	assert.Equal(1, hammingWeight(128))
	assert.Equal(31, hammingWeight(4294967293))
}
