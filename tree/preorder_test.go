package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreOrderRecursion(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ExceptedResOfPre, PreOrderRecursion(root))
}
