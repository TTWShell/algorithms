package lht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, numJewelsInStones("aA", "aAAbbbb"))
	assert.Equal(0, numJewelsInStones("z", "ZZ"))
}
