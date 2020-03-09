package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCut(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, minCut("aab"))
	assert.Equal(0, minCut("aa"))
}
