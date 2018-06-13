package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostOrderRecursion(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ExceptedResOfPost, PostOrderRecursion(root))
}

func TestPostOrderStack(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ExceptedResOfPost, PostOrderStack(root))
}
