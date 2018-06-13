package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numTrees(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, numTrees(3))
}
