package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countPrimeSetBits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, countPrimeSetBits(6, 10))
	assert.Equal(5, countPrimeSetBits(10, 15))
}
