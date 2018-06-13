package ldp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(1, uniquePaths(1, 2))
	assert.Equal(2, uniquePaths(2, 2))
	assert.Equal(193536720, uniquePaths(23, 12))
}
