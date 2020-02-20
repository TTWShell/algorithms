package ldp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDistance(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, minDistance("horse", "ros"))
	assert.Equal(5, minDistance("intention", "execution"))
}
