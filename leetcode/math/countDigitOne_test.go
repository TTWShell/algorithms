package lmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countDigitOne(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, countDigitOne(13))
}
