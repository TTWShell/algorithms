package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInOrderRecursion(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ExceptedResOfIn, InOrderRecursion(root))
}

func TestInOrderStack(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ExceptedResOfIn, InOrderStack(root))
}
