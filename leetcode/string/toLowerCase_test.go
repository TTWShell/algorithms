package lstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toLowerCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("hello", toLowerCase("Hello"))
	assert.Equal("here", toLowerCase("here"))
	assert.Equal("lovely", toLowerCase("LOVELY"))
}
